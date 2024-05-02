import { env } from 'node:process'
import fs from 'node:fs'

if (env.NITRO_PRESET === 'cloudflare-pages') {
  console.log('copying database to ./dist/database')
  fs.cpSync('./database/db.sqlite', './dist/database/db.sqlite', { recursive: true })
}
else {
  console.log('copying database to .output/database')
  fs.cpSync('./database/db.sqlite', '.output/database/db.sqlite', { recursive: true })
}
