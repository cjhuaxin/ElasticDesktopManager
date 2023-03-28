<template>
  <div style="height: 400px">
    <el-auto-resizer>
      <template #default="{ height, width }">
        <el-table-v2
          :columns="columns"
          :data="data"
          :width="width"
          :height="height"
          fixed
        />
      </template>
    </el-auto-resizer>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { models } from "../../wailsjs/go/models";
import { Search } from "../../wailsjs/go/service/Record";

const props = defineProps({
  indexId: String,
  indexName: String,
  connectionId: String,
  columns:Array<any>,
});


// const columns = props.columns
// const data = ref(new Array<any>());
//search data
let req = new models.QueryReq({
  connection_id: props.connectionId,
  index: props.indexName,
});
Search(req).then((result) => {
  for(let source of result.data.records.hits.hits){
    // data.value.push(source._source)
  }
  
  
});

const generateColumns = (length = 10, prefix = 'column-', props?: any) =>
  Array.from({ length }).map((_, columnIndex) => ({
    ...props,
    key: `${prefix}${columnIndex}`,
    dataKey: `${prefix}${columnIndex}`,
    title: `Column ${columnIndex}`,
    width: 150,
  }))
const generateData = (
  columns: ReturnType<typeof generateColumns>,
  length = 200,
  prefix = 'row-'
) =>
  Array.from({ length }).map((_, rowIndex) => {
    return columns.reduce(
      (rowData, column, columnIndex) => {
        rowData[column.dataKey] = `Row ${rowIndex} - Col ${columnIndex}`
        return rowData
      },
      {
        id: `${prefix}${rowIndex}`,
        parentId: null,
      }
    )
  })

  const columns = generateColumns(10)
  const data = generateData(columns, 1000)
</script>
