const esbuild = require("esbuild")

esbuild
 .build({
  entryPoints: ["frontend/Application.tsx", "frontend/css/styles.css"],
  bundle: true,
  sourcemap: true,
  minify: true,
  outdir: "public/assets",
  publicPath: './',
  loader: {
   '.png': 'file',
   '.svg': 'file',
   '.jpeg': 'file',
   '.jpg': 'file',
   '.ttf': 'file',
   '.otf': 'file',
  },
 })
 .then((r) => {
  console.log('✨ Ramses succeeded.');
 })