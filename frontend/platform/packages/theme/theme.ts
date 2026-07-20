import tokens from '../tokens/tokens.json';

export type ThemeMode = 'light' | 'dark' | 'high-contrast';

export function applyTheme(mode: ThemeMode) {
  const root = document.documentElement;
  
  // Clean classes
  root.classList.remove('dark', 'high-contrast');
  if (mode === 'dark') {
    root.classList.add('dark');
  } else if (mode === 'high-contrast') {
    root.classList.add('high-contrast');
  }

  const themeColors = tokens.themes[mode].colors;
  
  // Apply design tokens as CSS custom properties
  Object.entries(themeColors).forEach(([key, val]) => {
    root.style.setProperty(`--color-${key}`, val);
  });

  // Apply spacing
  Object.entries(tokens.spacing).forEach(([key, val]) => {
    root.style.setProperty(`--spacing-${key}`, val);
  });

  // Apply border radius
  Object.entries(tokens.radius).forEach(([key, val]) => {
    root.style.setProperty(`--radius-${key}`, val);
  });
}
