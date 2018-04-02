<!-- /views/admin/master.html -->

<!doctype html>

<html>
    <head>
        <meta charset="UTF-8">
        <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css"> 
        <title>{{.title}}</title> 
        
    </head>
    <body>
        <script src="//unpkg.com/vue/dist/vue.js"></script>
        <script src="//unpkg.com/element-ui/lib/index.js"></script>
        <div id="app">
            <el-container>
                <el-header>{{include "layouts/head"}}</el-header>
                <el-container>
                    <el-aside width="200px">{{template "ad" .}}</el-aside>
                    <el-container>
                    <el-main>{{template "content" .}}</el-main>
                    </el-container>
                </el-container>
                <el-footer>{{include "layouts/footer"}}</el-footer>
            </el-container>
        </div>
    </body> 
    <script>
        var Main = {
            data() {
                let item = {
                    date: '2016-05-02',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }
                return {
                    tableData: Array(20).fill(item)
                }
            }
        }
        var Ctor = Vue.extend(Main)
        new Ctor().$mount('#app')
    </script>

    <style type="text/css" scope>
        .el-header {
            background-color: #B3C0D1;
            color: #333;
            text-align: right;
        }

        .el-footer {
            background-color: #B3C0D1;
            color: #333;
            text-align: center;
            line-height: 60px;
        }
        
        .el-aside {
            background-color: #D3DCE6;
            color: #333;
            text-align: center;
            line-height: 200px;
            height: 100%;
        }
        
        .el-main {
            background-color: #E9EEF3;
            color: #333;
            text-align: center;
            line-height: 160px;
        }
        
        body > .el-container {
            margin-bottom: 40px;
        }
        
        .el-container:nth-child(5) .el-aside,
        .el-container:nth-child(6) .el-aside {
            line-height: 260px;
        }
        
        .el-container:nth-child(7) .el-aside {
            line-height: 320px;
        }

        .el-header {
            background-color: #B3C0D1;
            color: #333;
            line-height: 60px;
        } 

        .hello{ color: red;}
        hr{ border: 1px #ccc dashed;}
    </style>
</html>