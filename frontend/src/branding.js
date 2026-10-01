import brand from '../../branding/brand.json';

export default brand;

export function docsURL(page = 'index', anchor = '') {
  const normalized = page.replace(/\/$/, '') || 'index';
  return `${brand.docsURL}/${normalized}.md${anchor ? `#${anchor}` : ''}`;
}
