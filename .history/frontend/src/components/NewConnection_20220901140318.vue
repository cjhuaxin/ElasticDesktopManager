<template>
  <!-- Form -->
  <el-button text @click="dialogFormVisible = true">New Connection</el-button>

  <el-dialog v-model="dialogFormVisible" title="New Connection" draggable :close-on-click-modal="false">
    <el-form 
    label-width="100px"
    style="max-width: 450px"
    v-loading="loading"
    :label-position="'right'" 
    :rules="connectionRules" 
    :model="formNewConnection"
   >
      <el-form-item label="Name">
        <el-input v-model="formNewConnection.connectionName" autocomplete="off" />
      </el-form-item>
      <el-form-item label="Urls" prop="urls">
        <el-input v-model="formNewConnection.urls" autocomplete="off"
          placeholder="http://192.168.0.10:9200,http://192.168.0.11:9200" />
      </el-form-item>
      <el-form-item label="Username">
        <el-input v-model="formNewConnection.username" autocomplete="off" placeholder="optional" />
      </el-form-item>
      <el-form-item label="Password">
        <el-input v-model="formNewConnection.password" autocomplete="off" placeholder="optional" />
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="4">
          <div class="grid-content ep-bg-purple" />
          <el-button @click="testConn">Test Connection</el-button>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="24">
          <div id="testResult" class="test-result">{{ testConnectionData.result }}</div>
        </el-col>
      </el-row>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)">Create</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from "element-plus";
import { TestEsConnection } from "../../wailsjs/go/connection/Connection";
import { CreateEsConnection } from "../../wailsjs/go/connection/Connection";
import { models } from "../../wailsjs/go/models";

const dialogFormVisible = ref(false)
let loading = ref(false)
const ruleFormRef = ref<FormInstance>()
const formNewConnection = reactive({
  connectionName: '',
  urls: '',
  username: '',
  password: '',
})
const testConnectionData = reactive({
  result: "",
})

const connectionRules = reactive<FormRules>({
  urls: [
    { required: true, message: 'Please input urls', trigger: 'blur' },
  ],

})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) {
    console.log("invalid formEl")
    return
  }
  await formEl.validate((valid, fields) => {
    if (valid) {
      let req = new models.NewConnectionReq({
        "name": formNewConnection.connectionName,
        "urls": formNewConnection.urls,
        "username": formNewConnection.username,
        "password": formNewConnection.password
      })
      CreateEsConnection(req).then(result => {
        if (result == "") {
          testConnectionData.result = "Succeeded"
        } else {
          testConnectionData.result = "Failed: " + result
        }
      })
    } else {
      testConnectionData.result = "Please "
    }
  })
}

function testConn() {
  loading = ref(true)
  let req = new models.NewConnectionReq({
    "name": formNewConnection.connectionName,
    "urls": formNewConnection.urls,
    "username": formNewConnection.username,
    "password": formNewConnection.password
  })

  TestEsConnection(req).then(result => {
    loading = ref(true)
    if (result == "") {
      testConnectionData.result = "Succeeded"
    } else {
      testConnectionData.result = "Failed: " + result
    }
  })
}
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}

.test-result {
  text-align: start;
}
</style>