package jsonrefs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Node interface {
	json.Marshaler
	GetPathPart(p string) (Node, error)
}

type ObjectNode struct {
	Map map[string]Node
}

type ArrayNode struct {
	Items []Node
}

type LeafNode struct {
	Raw json.RawMessage
}

type RefNode struct {
	Ref string
}

func Replace(raw *json.RawMessage) (*json.RawMessage, error) {
	if raw == nil {
		return nil, errors.New("json raw message is nil")
	}

	root, err := unmarshalNode(*raw)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json as node: %w", err)
	}

	resolved, err := resolveNode(root, root, nil)
	if err != nil {
		return nil, fmt.Errorf("replace json refs: %w", err)
	}

	resolvedBytes, err := json.Marshal(resolved)
	if err != nil {
		return nil, fmt.Errorf("marshal resolved node: %w", err)
	}

	resolvedRaw := json.RawMessage(resolvedBytes)
	return &resolvedRaw, nil
}

func unmarshalNode(raw json.RawMessage) (Node, error) {
	var node nodeJSON
	if err := json.Unmarshal(raw, &node); err != nil {
		return nil, err
	}
	if node.Node == nil {
		return nil, errors.New("node is nil")
	}
	return node.Node, nil
}

type nodeJSON struct {
	Node Node
}

func (n *nodeJSON) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return errors.New("empty json")
	}

	switch trimmed[0] {
	case '{':
		var rawMap map[string]json.RawMessage
		if err := json.Unmarshal(data, &rawMap); err != nil {
			return fmt.Errorf("unmarshal object: %w", err)
		}

		if _, ok := rawMap["$ref"]; ok {
			var ref RefNode
			if err := json.Unmarshal(data, &ref); err != nil {
				return err
			}
			n.Node = &ref
			return nil
		}

		var object ObjectNode
		if err := json.Unmarshal(data, &object); err != nil {
			return err
		}
		n.Node = &object
		return nil
	case '[':
		var array ArrayNode
		if err := json.Unmarshal(data, &array); err != nil {
			return err
		}
		n.Node = &array
		return nil
	default:
		var leaf LeafNode
		if err := json.Unmarshal(data, &leaf); err != nil {
			return err
		}
		n.Node = &leaf
		return nil
	}
}

func (n nodeJSON) MarshalJSON() ([]byte, error) {
	if n.Node == nil {
		return nil, errors.New("node is nil")
	}
	return json.Marshal(n.Node)
}

func (n *ObjectNode) UnmarshalJSON(data []byte) error {
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return fmt.Errorf("unmarshal object: %w", err)
	}

	n.Map = make(map[string]Node, len(rawMap))
	for key, rawValue := range rawMap {
		child, err := unmarshalNode(rawValue)
		if err != nil {
			return fmt.Errorf("unmarshal object key %q: %w", key, err)
		}
		n.Map[key] = child
	}

	return nil
}

func (n ObjectNode) MarshalJSON() ([]byte, error) {
	rawMap := make(map[string]json.RawMessage, len(n.Map))
	for key, child := range n.Map {
		if child == nil {
			return nil, fmt.Errorf("marshal object key %q: node is nil", key)
		}

		rawValue, err := json.Marshal(child)
		if err != nil {
			return nil, fmt.Errorf("marshal object key %q: %w", key, err)
		}
		rawMap[key] = rawValue
	}

	return json.Marshal(rawMap)
}

func (n ObjectNode) GetPathPart(p string) (Node, error) {
	child, ok := n.Map[p]
	if !ok {
		return nil, fmt.Errorf("path part %q not found", p)
	}
	return child, nil
}

func (n *ArrayNode) UnmarshalJSON(data []byte) error {
	var rawItems []json.RawMessage
	if err := json.Unmarshal(data, &rawItems); err != nil {
		return fmt.Errorf("unmarshal array: %w", err)
	}

	n.Items = make([]Node, len(rawItems))
	for i, rawItem := range rawItems {
		child, err := unmarshalNode(rawItem)
		if err != nil {
			return fmt.Errorf("unmarshal array index %d: %w", i, err)
		}
		n.Items[i] = child
	}

	return nil
}

func (n ArrayNode) MarshalJSON() ([]byte, error) {
	rawItems := make([]json.RawMessage, len(n.Items))
	for i, child := range n.Items {
		if child == nil {
			return nil, fmt.Errorf("marshal array index %d: node is nil", i)
		}

		rawItem, err := json.Marshal(child)
		if err != nil {
			return nil, fmt.Errorf("marshal array index %d: %w", i, err)
		}
		rawItems[i] = rawItem
	}

	return json.Marshal(rawItems)
}

func (n ArrayNode) GetPathPart(p string) (Node, error) {
	return nil, fmt.Errorf("cannot get path part %q from array", p)
}

func (n *LeafNode) UnmarshalJSON(data []byte) error {
	n.Raw = append(json.RawMessage(nil), data...)
	return nil
}

func (n LeafNode) MarshalJSON() ([]byte, error) {
	if len(n.Raw) == 0 {
		return []byte("null"), nil
	}
	return append([]byte(nil), n.Raw...), nil
}

func (n LeafNode) GetPathPart(p string) (Node, error) {
	return nil, fmt.Errorf("cannot get path part %q from leaf", p)
}

func (n *RefNode) UnmarshalJSON(data []byte) error {
	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return fmt.Errorf("unmarshal ref object: %w", err)
	}

	rawRef, ok := rawMap["$ref"]
	if !ok {
		return errors.New("ref object does not contain $ref")
	}

	if err := json.Unmarshal(rawRef, &n.Ref); err != nil {
		return fmt.Errorf("unmarshal $ref string: %w", err)
	}

	return nil
}

func (n RefNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"$ref": n.Ref})
}

func (n RefNode) GetPathPart(p string) (Node, error) {
	return nil, fmt.Errorf("cannot get path part %q from ref %q", p, n.Ref)
}

func resolveNode(root Node, node Node, stack []string) (Node, error) {
	switch n := node.(type) {
	case *RefNode:
		return resolveRef(root, n.Ref, stack)
	case RefNode:
		return resolveRef(root, n.Ref, stack)
	case *ObjectNode:
		for key, child := range n.Map {
			resolved, err := resolveNode(root, child, stack)
			if err != nil {
				return nil, fmt.Errorf("resolve object key %q: %w", key, err)
			}
			n.Map[key] = resolved
		}
		return n, nil
	case ObjectNode:
		for key, child := range n.Map {
			resolved, err := resolveNode(root, child, stack)
			if err != nil {
				return nil, fmt.Errorf("resolve object key %q: %w", key, err)
			}
			n.Map[key] = resolved
		}
		return n, nil
	case *ArrayNode:
		for i, child := range n.Items {
			resolved, err := resolveNode(root, child, stack)
			if err != nil {
				return nil, fmt.Errorf("resolve array index %d: %w", i, err)
			}
			n.Items[i] = resolved
		}
		return n, nil
	case ArrayNode:
		for i, child := range n.Items {
			resolved, err := resolveNode(root, child, stack)
			if err != nil {
				return nil, fmt.Errorf("resolve array index %d: %w", i, err)
			}
			n.Items[i] = resolved
		}
		return n, nil
	case *LeafNode, LeafNode:
		return n, nil
	default:
		return nil, fmt.Errorf("unknown node type %T", node)
	}
}

func resolveRef(root Node, ref string, stack []string) (Node, error) {
	for _, seen := range stack {
		if seen == ref {
			return nil, fmt.Errorf("reference cycle for %q", ref)
		}
	}

	target, err := refTarget(root, ref)
	if err != nil {
		return nil, err
	}

	return resolveNode(root, target, append(stack, ref))
}

func refTarget(root Node, ref string) (Node, error) {
	if !strings.HasPrefix(ref, "#/") {
		return nil, fmt.Errorf("$ref %q is invalid: must start with #/", ref)
	}

	parsed, err := url.Parse(ref)
	if err != nil {
		return nil, fmt.Errorf("parse $ref %q: %w", ref, err)
	}

	if !strings.HasPrefix(parsed.Fragment, "/") {
		return nil, fmt.Errorf("$ref %q is invalid: fragment must start with /", ref)
	}

	node := root
	for _, rawPart := range strings.Split(parsed.Fragment[1:], "/") {
		part, err := unescapePathPart(rawPart)
		if err != nil {
			return nil, fmt.Errorf("unescape $ref %q path part %q: %w", ref, rawPart, err)
		}

		node, err = node.GetPathPart(part)
		if err != nil {
			return nil, fmt.Errorf("get $ref %q path part %q: %w", ref, part, err)
		}
	}

	return node, nil
}

func unescapePathPart(part string) (string, error) {
	unescaped := make([]byte, 0, len(part))

	for i := 0; i < len(part); i++ {
		if part[i] != '~' {
			unescaped = append(unescaped, part[i])
			continue
		}

		if i+1 >= len(part) {
			return "", errors.New("~ must be followed by 0 or 1")
		}

		switch part[i+1] {
		case '0':
			unescaped = append(unescaped, '~')
		case '1':
			unescaped = append(unescaped, '/')
		default:
			return "", fmt.Errorf("~%c is invalid", part[i+1])
		}
		i++
	}

	return string(unescaped), nil
}
