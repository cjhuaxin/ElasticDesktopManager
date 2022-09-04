<script setup>
import NewConnection from './NewConnection.vue'
</script>

<template>
  <el-row :gutter="20">
    <el-col :span="4"><div class="grid-content ep-bg-purple" />
      <NewConnection/>
    </el-col>
  </el-row>
</template>

<style>
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>
