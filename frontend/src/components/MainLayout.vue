<script setup>
import NewConnection from './NewConnection.vue'
import ConnectionList from './ConnectionList.vue'
</script>

<template>
  <el-row>
    <el-col :span="4">
      <div class="grid-content ep-bg-purple" />
      <NewConnection/>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="8">
      <div class="grid-content ep-bg-purple" />
      <ConnectionList/>
    </el-col>
  </el-row>
</template>

<style>
.grid-content {
  border-radius: 4px;
  min-height: 12px;
}
</style>
