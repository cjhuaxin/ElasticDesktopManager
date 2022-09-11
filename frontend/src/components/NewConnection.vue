<template>
  <!-- Form -->
  <el-button @click="dialogFormVisible = true" type="info">New Connection</el-button>
  <el-icon><Refresh /></el-icon>


  <el-dialog v-model="dialogFormVisible" title="New Connection" draggable :close-on-click-modal="false">
    <el-form ref="ruleFormRef" label-width="100px" style="max-width: 450px" v-loading="loading" :label-position="'top'"
      :rules="connectionRules" :model="formNewConnection">
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
          <div :class="testErr ? 'test-failed' : 'test-success'">{{ testConnectionData.result }}</div>
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
import { FormInstance, FormRules, ElNotification } from "element-plus";
import { Refresh } from '@element-plus/icons'
import { TestEsConnection, CreateEsConnection } from "../../wailsjs/go/service/Connection";
import { models } from "../../wailsjs/go/models";
import bus from './mitt'
import { ref, reactive } from "vue";

const testErr = ref(false)
const dialogFormVisible = ref(false)
const loading = ref(false)
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

function testConn() {
  loading.value = true
  let req = new models.NewConnectionReq({
    "name": formNewConnection.connectionName,
    "urls": formNewConnection.urls,
    "username": formNewConnection.username,
    "password": formNewConnection.password
  })

  TestEsConnection(req).then(result => {
    loading.value = false
    if (result.err_msg == "") {
      testErr.value = false
      testConnectionData.result = "Succeeded"
    } else {
      testErr.value = true
      testConnectionData.result = "Failed: " + result.err_msg
    }
  })
}

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) {
    console.log("invalid formEl")
    return
  }
  await formEl.validate((valid, fields) => {
    if (valid) {
      loading.value = true
      let req = new models.NewConnectionReq({
        "name": formNewConnection.connectionName,
        "urls": formNewConnection.urls,
        "username": formNewConnection.username,
        "password": formNewConnection.password
      })
      CreateEsConnection(req).then(result => {
        loading.value = false
        dialogFormVisible.value = false
        bus.emit('handleAddNewConnection')

        if (result.err_msg == "") {
          ElNotification.success({
            title: 'Success',
            message: 'Create Connection Success',
            showClose: false,
          })
        } else {
          ElNotification.error({
            title: 'Failed',
            message: result.err_msg,
            showClose: false,
          })
        }
      })
    } else {
      testConnectionData.result = "Please input the required fields"
    }
  })
}
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}

.test-success {
  color: rgb(32, 97, 49);
  text-align: start;
}

.test-failed {
  color: rgb(190, 85, 81);
  text-align: start;
}
</style>