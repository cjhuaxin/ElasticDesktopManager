<script setup>
import NewConnection from './NewConnection.vue'
import ConnectionList from './ConnectionList.vue'
import DataTab from './DataTab.vue'
</script>

<template>
  <el-row>
    <el-col :xs="4" :sm="4" :md="4" :lg="4" :xl="4">
      <div class="grid-content ep-bg-purple" />
      <NewConnection />
    </el-col>
    <el-col :xs="20" :sm="20" :md="16" :lg="16" :xl="16">
      <div class="grid-content ep-bg-purple" />
      <DataTab />
    </el-col>
  </el-row>
  <el-row>
    <el-col :xs="4" :sm="4" :md="8" :lg="8" :xl="8">
      <div class="grid-content ep-bg-purple" />
      <ConnectionList />
    </el-col>
  </el-row>
</template>

<style>
.grid-content {
  border-radius: 4px;
  min-height: 12px;
}
</style>
