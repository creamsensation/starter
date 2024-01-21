import esbuild from 'esbuild'
import postcss from 'esbuild-postcss'
import manifest from 'esbuild-plugin-manifest'
import { clean } from 'esbuild-plugin-clean'
import { watch } from 'chokidar'
import prettyBytes from 'pretty-bytes'
import chalk from 'chalk'

import { assetsConfig } from './assets.config.ts'

function build() {
	console.clear()
	console.log('Build started...')
	esbuild
		.build({
			entryPoints: ['./styles/main.css', './scripts/main.ts'],
			entryNames: '[name]-[hash]',
			bundle: true,
			sourcemap: true,
			minify: false,
			outdir: `./${assetsConfig.publicPath}/${assetsConfig.outputPath}`,
			format: 'esm',
			plugins: [
				postcss(),
				clean({
					patterns: [`./${assetsConfig.publicPath}/${assetsConfig.outputPath}/*`],
				}),
				manifest(),
			],
		})
		.then(v => {
			if (typeof v.metafile !== 'undefined') {
				for (const k in v.metafile.outputs) {
					if (k.endsWith('.map')) {
						continue
					}
					console.log(`${chalk.green(k)}: ${prettyBytes(v.metafile.outputs[k].bytes)}`)
				}
			}
			console.log('Build completed')
			console.log('Watching...')
		})
		.catch(() => process.exit(1))
}

build()

watch(['./styles/**/*.css', './scripts/**/*.ts', '../app/**/*.go']).on('change', () => {
	build()
})
