<template>
  <el-tree :props="defaultProps" :load="loadNode" lazy />
</template>

<script lang="ts" setup>
import Node from 'element-plus/es/components/tree/src/model/node'
import { GetSavedConnectionList } from "../../wailsjs/go/service/Connection";

import bus from './mitt'

interface Tree {
  id: string
  name: string
  leaf?: boolean
}
interface ConnectionInfo {
  id: string;
  name: string;
}

const defaultProps = {
  label: 'name',
  children: 'zones',
  isLeaf: 'leaf',
}

let rootNode: Node
let rootResolve: (data: Tree[]) => void

const loadNode = (node: Node, resolve: (data: Tree[]) => void) => {
  console.log(node)
  if (node.level != 0) {
    return resolve([])
  }

  switch (node.level) {
    case 0:
      //reload the whole tree
      rootNode = node
      rootResolve = resolve
      return getSavedConnectionList(resolve)
    case 1:
      break;
  }


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
    