const fs = require('node:fs');
const path = require('node:path');

const root = path.resolve(__dirname, '..');
fs.copyFileSync(
  path.join(root, 'frontend/node_modules/altcha/dist/altcha.umd.cjs'),
  path.join(root, 'static/public/static/altcha.umd.js'),
);
