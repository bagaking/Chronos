{{define "content"}}
    <el-table :data="tableData">
        <el-table-column prop="date" label="日期" width="140">
        </el-table-column>
        <el-table-column prop="name" label="姓名" width="120">
        </el-table-column>
        <el-table-column prop="address" label="地址">
        </el-table-column>

    </el-table>

    <el-button @click="visible = true">按钮</el-button>
    <el-dialog :visible.sync="visible" title="Hello world">
        <p>欢迎使用 Element</p>
    </el-dialog>
    <h1 class="hello">This is content!!!!</h1>
    <el-row>
        <el-col :span="12">123 + 333 = {{call $.add 123 333}}</el-col>
        <el-col :span="12"> 
            <el-dropdown>
                <i class="el-icon-setting" style="margin-right: 15px"></i>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>Run</el-dropdown-item>
                    <el-dropdown-item>Edit</el-dropdown-item>
                    <el-dropdown-item>Remove</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown> 
        </el-col>
    </el-row>
    <el-row>
        <el-col :span="12">123 + 333 = {{call $.add 123 333}}</el-col>
        <el-col :span="12"> 
            <el-dropdown>
                <i class="el-icon-setting" style="margin-right: 15px"></i>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>Run</el-dropdown-item>
                    <el-dropdown-item>Edit</el-dropdown-item>
                    <el-dropdown-item>Remove</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown> 
        </el-col>
    </el-row>
    
    <p>123 - 100= {{sub 123 100}}</p>
    <hr>
    <p><a href="/page">Page render</a></p>
{{end}}


    