<script setup>
import NewConnection from "./NewConnection.vue";
import ConnectionList from "./ConnectionList.vue";
import DataTab from "./DataTab.vue";
</script>

<template>
  <el-row>
    <el-col :xs="3" :sm="3" :md="7" :lg="7" :xl="7">
      <div class="grid-content ep-bg-purple">
        <NewConnection />
      </div>
      <div class="grid-content ep-bg-purple">
        <ConnectionList />
      </div>
    </el-col>
    <el-col :xs="21" :sm="21" :md="17" :lg="17" :xl="17">
      <div class="grid-content ep-bg-purple" />
      <DataTab />
    </el-col>
  </el-row>
</template>

<style>
.grid-content {
  border-radius: 4px;
  min-height: 12px;
}
</style>
