<template>
  <el-tree
    :props="defaultProps"
    :load="loadNode"
    @node-click="handleIndexClick"
    indent="10"
    highlight-current
    lazy
  >
    <template #default="{ node, data }">
      <i style="health-icon" v-if="node.isLeaf">
        <svg height="14" width="14">
          <circle cx="7" cy="9" r="4" :class="data.healthStyle"></circle>
        </svg>
      </i>
      <span>
        <span>{{ data.label }}</span>
        <slot :data="data" :node="node" />
      </span>
    </template>
  </el-tree>
</template>

<script lang="ts" setup>
import Node from "element-plus/es/components/tree/src/model/node";
import { GetSavedConnectionList } from "../../wailsjs/go/service/Connection";
import { CatIndex } from "../../wailsjs/go/service/Index";
import { models } from "../../wailsjs/go/models";
import { ElNotification } from "element-plus";
import emitter from "@/utils/emitter";

interface Tree {
  id: string;
  connectionId: string;
  label: string;
  name: string;
  leaf?: boolean;
  healthStyle: string;
}

interface ConnectionInfo {
  id: string;
  name: string;
}

interface IndexInfo {
  uuid: string;
  index: string;
  health: string;
  docs_count: number;
  store_size: string;
}

const defaultProps = {
  label: "name",
  children: "children",
  isLeaf: "leaf",
};

let rootNode: Node;
let rootResolve: (data: Tree[]) => void;

const loadNode = (node: Node, resolve: (data: Tree[]) => void) => {
  switch (node.level) {
    case 0:
      //reload the whole tree
      rootNode = node;
      rootResolve = resolve;
      return getSavedConnectionList(resolve);
    case 1:
      return listIndex(node, resolve);
  }
};

const handleIndexClick = (data: Node) => {
  emitter.emit("index-click", data);
};

function listIndex(node: Node, resolve: (data: Tree[]) => void) {
  let connectionId = node.data.id;
  let req = new models.CatIndexReq({
    id: connectionId,
  });
  let treeList: Array<Tree> = new Array<Tree>();
  CatIndex(req).then((result) => {
    if (result.err_msg != "") {
      ElNotification.error({
        title: "Failed",
        message: result.err_msg,
        showClose: false,
      });
      return resolve(treeList);
    }

    let resData = result.data as Array<IndexInfo>;
    for (let data of resData) {
      let treeNode: Tree = {
        id: data.uuid,
        connectionId: connectionId,
        label: `${data.index} [${data.docs_count}]`,
        // label: `${data.index} [${data.docs_count}/${data.store_size}]`,
        name: data.index,
        leaf: true,
        healthStyle: `health-${data.health}`,
      };
      treeList.push(treeNode);
    }
    return resolve(treeList);
  });
}

function getSavedConnectionList(resolve: (data: Tree[]) => void) {
  let treeList: Array<Tree> = new Array<Tree>();
  GetSavedConnectionList().then((result) => {
    let resData = result.data as Array<ConnectionInfo>;
    for (let data of resData) {
      let treeNode: Tree = {
        id: data.id,
        connectionId: data.id,
        label: data.name,
        name: data.name,
        leaf: false,
        healthStyle: "",
      };
      treeList.push(treeNode);
    }

    return resolve(treeList);
  });
}

function resetNode() {
  const theChildren = rootNode.childNodes;
  theChildren.splice(0, theChildren.length);
  loadNode(rootNode, rootResolve);
}
emitter.on("add-new-connection", resetNode);
</script>

<style>
.health-icon {
  margin-right: 1%;
}

.el-tree-node__expand-icon.is-leaf
{
    display: none;
}
</style>
