<template>
  <div style="height: 500px">
    <el-auto-resizer>
      <template #default="{ width, height }">
        <el-table
          :data="records"
          :width="width"
          :height="height"
          :table-layout="auto"
          stripe
        >
          <el-table-column
            v-for="item in columns"
            :key="item.key"
            :prop="item.dataKey"
            :label="item.title"
          >
          </el-table-column>
        </el-table>
      </template>
    </el-auto-resizer>
  </div>
  <el-pagination
    background
    layout="prev, next, sizes"
    :page-sizes="[10, 20, 50, 100]"
    :total="1000"
  />
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { models } from "../../wailsjs/go/models";
import { Search } from "../../wailsjs/go/service/Record";
import type { Column } from "element-plus";
import { auto } from "@popperjs/core";

const props = defineProps({
  indexId: String,
  indexName: String,
  connectionId: String,
});

const columns = ref(new Array<Column<any>>());
const records = ref([new Array<any>()]);
//search data
let req = new models.QueryReq({
  connection_id: props.connectionId,
  index: props.indexName,
});
Search(req).then((result) => {
  let properties = new Array<Column<any>>();
  let sources = new Array<any>();
  for (let source of result.data.records.hits.hits) {
    let keys = Object.keys(source._source);
    if (properties.length == 0 || properties.length < keys.length) {
      for (let key of keys) {
        properties.push({
          key: key,
          title: key,
          dataKey: key,
          width: 0,
        });
      }
    }
    sources.push(source._source);
  }
  columns.value = properties;
  records.value = sources;
});
</script>
