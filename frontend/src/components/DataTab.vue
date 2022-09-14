<template>
    <el-tabs v-model="activeTabName" type="card" class="demo-tabs" closable v-if=showDataTab>
        <el-tab-pane label="User" name="first">
            <RecordTab />
        </el-tab-pane>
    </el-tabs>
</template>
<script lang="ts" setup>
import Node from 'element-plus/es/components/tree/src/model/node'
import { ref } from 'vue'
import emitter from "@/utils/emitter";
import RecordTab from './RecordTab.vue'

const showDataTab = ref(false)
const activeTabName = ref('')
const dataTabs = ref([])

const clickIndex = function (param: any) {
    console.log(param)
    if (param.leaf) {
        console.log(param)
        showDataTab.value = true
    }
}

emitter.on('index-click', clickIndex)
</script>