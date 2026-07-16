import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

export default defineConfig({
  site: 'https://djosh34.github.io',
  base: '/openapi-validate',
  integrations: [
    starlight({
      title: 'openapi-validate',
      sidebar: [
        { label: 'Getting started', link: '/' },
        { label: 'Philosophy', link: '/philosophy/' },
        { label: 'Query decoding', link: '/query-decoding/' },
        { label: 'Architecture', link: '/architecture/' },
        { label: 'Roadmap', link: '/roadmap/' },
      ],
    }),
  ],
});
