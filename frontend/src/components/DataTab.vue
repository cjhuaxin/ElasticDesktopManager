<template>
  <el-tabs
    v-model="activeTabName"
    type="border-card"
    @tab-remove="removeTab"
    closable
    v-if="showDataTab"
  >
    <el-tab-pane
      v-for="item in dataTabs"
      :key="item.name"
      :label="item.title"
      :name="item.name"
    >
      <Record
        :indexId="item.indexId"
        :indexName="item.name"
        :connectionId="item.connectionId"
        :columns="item.columns"
      />
    </el-tab-pane>
  </el-tabs>
</template>
<script lang="ts" setup>
import { ref, shallowRef } from "vue";
import emitter from "@/utils/emitter";
import Record from "./RecordTab.vue";
import { models } from "../../wailsjs/go/models";
import { GetProperties } from "../../wailsjs/go/service/Record";

interface Tab {
  indexId: string;
  connectionId: string;
  name: string;
  title: string;
  columns: Array<any>;
  content: any;
}

const showDataTab = ref(false);
const activeTabName = ref("");
const dataTabs = ref(new Array<Tab>());

const clickIndex = function (param: any) {
  if (param.leaf) {
    showDataTab.value = true;
    let req = new models.QueryReq({
      connection_id: param.connectionId,
      index: param.name,
    });
    GetProperties(req).then((result) => {
      let newTab: Tab = {
        indexId: param.id,
        connectionId: param.connectionId,
        name: param.name,
        title: param.name,
        columns: result.data,
        content: shallowRef(Record),
      };
      dataTabs.value.push(newTab);
      activeTabName.value = param.name;
    });
  }
};

const removeTab = (targetName: string) => {
  console.log(targetName);
  const tabs = dataTabs.value;
  let activeName = activeTabName.value;
  if (activeName === targetName) {
    tabs.forEach((tab, index) => {
      if (tab.name === targetName) {
        const nextTab = tabs[index + 1] || tabs[index - 1];
        console.log("nextTab" + nextTab);
        if (nextTab) {
          activeName = nextTab.name;
        }
      }
    });
  }
  activeTabName.value = activeName;
  dataTabs.value = tabs.filter((tab) => tab.name !== targetName);
};

emitter.on("index-click", clickIndex);
</script>
