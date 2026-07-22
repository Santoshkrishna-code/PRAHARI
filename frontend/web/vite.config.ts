import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

export default defineConfig({
  base: './',
  plugins: [react()],
  resolve: {
    alias: {
      '@prahari/tokens': path.resolve(__dirname, '../platform/packages/tokens'),
      '@prahari/theme': path.resolve(__dirname, '../platform/packages/theme'),
      '@prahari/components': path.resolve(__dirname, '../platform/packages/components'),
      '@prahari/layouts': path.resolve(__dirname, '../platform/packages/layouts'),
      '@prahari/navigation': path.resolve(__dirname, '../platform/packages/navigation'),
      '@prahari/charts': path.resolve(__dirname, '../platform/packages/charts'),
      '@prahari/utils': path.resolve(__dirname, '../platform/packages/utils'),
      '@prahari/state': path.resolve(__dirname, '../platform/packages/state')
    }
  }
});
