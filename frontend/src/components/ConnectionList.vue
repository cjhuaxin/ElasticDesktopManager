<template>
  <el-tree :props="defaultProps" :load="loadNode" lazy />
</template>

<script lang="ts" setup>
import Node from 'element-plus/es/components/tree/src/model/node'
import { GetSavedConnectionList } from "../../wailsjs/go/service/Connection";
import { CatIndex } from "../../wailsjs/go/service/Index";
import { models } from "../../wailsjs/go/models";

import bus from './mitt'
import { capitalize } from 'lodash';

interface Tree {
  id: string
  name: string
  leaf?: boolean
  icon: string
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
  label: 'name',
  children: 'index',
  isLeaf: 'leaf',
}

let rootNode: Node
let rootResolve: (data: Tree[]) => void

const loadNode = (node: Node, resolve: (data: Tree[]) => void) => {
  switch (node.level) {
    case 0:
      //reload the whole tree
      rootNode = node
      rootResolve = resolve
      return getSavedConnectionList(resolve)
    case 1:
      return listIndex(node.data.id, resolve)
  }


}

function listIndex(connectionId: string, resolve: (data: Tree[]) => void) {
  let req = new models.CatIndexReq({
    id: connectionId
  })
  let treeList: Array<Tree> = new Array<Tree>()
  CatIndex(req).then(result => {
    let resData = result.data as Array<IndexInfo>
    for (let data of resData) {
      let treeNode: Tree = {
        id: data.uuid,
        name: `${data.index} [${data.health}/${data.docs_count}/${data.store_size}]`,
        leaf: true,
        icon: ""
      }
      treeList.push(treeNode)
    }
    return resolve(treeList)
  })

}

function getSavedConnectionList(resolve: (data: Tree[]) => void) {
  let treeList: Array<Tree> = new Array<Tree>()
  GetSavedConnectionList().then(result => {
    let resData = result.data as Array<ConnectionInfo>
    for (let data of resData) {
      let treeNode: Tree = {
        id: data.id,
        name: data.name
      }
      treeList.push(treeNode)
    }


    return resolve(treeList)
  })

}

function resetNode() {
  var theChildren = rootNode.childNodes
  theChildren.splice(0, theChildren.length)
  loadNode(rootNode, rootResolve)
}

bus.on('handleAddNewConnetion', resetNode)
</script>
    