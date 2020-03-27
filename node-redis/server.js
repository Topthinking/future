const Koa = require('koa')
const redis = require('redis')
const os = require('os')

const client = redis.createClient({
    host: process.env.REDIS_HOST,
    port: 6379
})

const app = new Koa()

const run = async (cb) => new Promise((resolve)=>{
    cb((err,reply)=>{
        resolve(reply)
    })
})

app.use(async(ctx) => {
    await run((cb)=>client.incr('hits',cb))
    const num = await run((cb) => client.get('hits',cb))
	ctx.body = "当前访问次数：" + num + "，当前主机名称：" + os.hostname() + "\n"
})

app.listen(80)