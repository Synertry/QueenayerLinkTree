/*
 *           QueenayerLinkTree
 *     Copyright (c) Synertry 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vite.dev/config/
export default defineConfig({
  // Set root directory to web folder
  root: '.',
  // Use root path for all asset URLs
  base: '/',
  plugins: [svelte()],
  build: {
    // Output directory will be 'dist' within the web folder
    outDir: 'dist',
    // Generate a manifest for potential asset fingerprinting
    manifest: true,

    target: 'esnext',
    cssTarget: 'esnext',
    cssMinify: "lightningcss"
  },

  css: {
    transformer: "lightningcss"
  }
})
