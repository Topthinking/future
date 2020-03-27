const Koa = require('koa')
const mysql = require('mysql');

const connection = mysql.createConnection({
  host     : process.env.MYSQL_HOST,
  user     : 'root',
  password : 'root',
  database : 'mysql'
});

const app = new Koa()

const run = (sql) => new Promise((resolve)=>{    
    connection.query(sql, function (error, results, fields) {
        if (error) throw error;
        resolve(results);
    });    
})

app.use(async(ctx)=>{
    const data = await run("select * from user")
	ctx.body = "hello world" + data[0].User
})

app.listen(5000)