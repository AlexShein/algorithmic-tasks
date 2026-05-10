test = require("./test")

/*Codewars description:

Goal
Your goal in this Kata is to complete the function topoSort(nodes) which will return a valid topological sort for a given DAG, unless you can spot a cycle in the input nodes. In that case, you need to throw an error with a specific message (see below).

The input can ba rather big, with up to hundreds of nodes, that can be linked to hundreds of other nodes, so watch out for performances.

The DAG is provided to you as a collection of Node objects. A Node is defined as follows:

function Node(id, i, o) {
  this.id = id;
  this.in = i;
  this.out = o;
}
id is a unique identifier for each node. It can be either numbers or string and there are no assumptions that can be done about the values.
in is an array containing references to all the Nodes in the graph which point to this Node.
out is an array containing references to all the Nodes which this Node points to.
If a Node has no inbound edges in will be an empty array. If a Node has no outbound edges out will be an empty array.
Node instances are frozen: you cannot modify them.
Returned Value
The return value of topoSort(nodes) should be an array of all Nodes from the input in Topological Sort order.

Error Condition
In addition to returning a correct Topological Sort the topoSort function must also catch a particular error condition: a cycle in the graph. If topoSort discovers a cycle in the graph it must throw an error. A string will work too, but in any case, the message has to be the following string:

// cycle found

*/

// Returns an array of nodes which are sorted in topological order.
// Uses the Kahn's algorithm,
// performance is suboptimal due to iterating through each vertex's out / in arrays
function topoSort(nodes) {
  const sortedNodes = new Array()
  const stack = new Array()
  const visitedNodes = new Set()
  const allNodeIds = new Set(nodes.map((node) => node.id))

  for (const node of nodes) {
    if (node.in.length == 0) {
      stack.push(node)
    }
  }

  while (stack.length > 0) {
    node = stack.pop()
    visitedNodes.add(node.id)
    sortedNodes.push(node)
    for (const nextNode of node.out) {
      if (nextNode.in.map((inNode) => visitedNodes.has(inNode.id)).every(Boolean)) {
        stack.push(nextNode)
      }
    }
  }

  if (sortedNodes.length !== nodes.length) {
    throw new Error("Graph has a cycle")
  }
  return sortedNodes
}

// Returns an array of nodes which are sorted in topological order.
// Uses the DFS algorithm
function topoSort(nodes) {
  const processedNodeIds = new Set()
  const tempMarkNodeIds = new Set()
  const sortedNodes = new Array()
  for (const node of nodes) {
    if (!processedNodeIds.has(node.id)) visitNode(node)
  }

  function visitNode(node) {
    if (processedNodeIds.has(node.id)) return
    if (tempMarkNodeIds.has(node.id)) throw new Error("Graph has a cycle")

    tempMarkNodeIds.add(node.id)
    for (m of node.out) visitNode(m)
    processedNodeIds.add(node.id)
    sortedNodes.unshift(node)
  }
  return sortedNodes
}
