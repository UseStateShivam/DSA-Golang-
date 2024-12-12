# Binary Tree Inorder Traversal

## Key learnings
 
- Binary Tree: A binary tree is a hierarchical data structure in which each node has at most two children, referred to as the left child and the right child. Binary trees are foundational structures in computer science, commonly used in tasks like searching, sorting, and hierarchical data representation.
- **Node**: A single element in the tree containing data.
- **Root**: The topmost node of the binary tree.
- **Child**: Nodes extending from another node.
- **Parent**: A node that has child nodes.
- **Leaf**: A node with no children.
- **Height**: The length of the longest path from the root to a leaf.
- **Depth**: The level of a node, i.e., the number of edges from the root to the node.
- **The order of binary trees are of the following types**: 
    For the binary tree:
    ``` lua
                2
            3       4
        5
    ```
    In-order DFS: 5 3 2 4 
    Pre-order DFS: 2 3 5 4 
    Post-order DFS: 5 3 4 2 
    Level order: 2 3 4 5 
- In-order traversal follows the order Left → Root → Right at each node recursively:
    **Start at the root**:
        Recursively visit the left subtree of the root.
    **Visit the root (or the current node)**:
        Process the value of the current node (e.g., print it or add it to a result list).
    **Traverse the right subtree**:
        Recursively visit the right subtree of the current node.
- For the given tree:
    ``` lua
            A
           / \
          B   C
         / \
        D   E
    ```
    **In-Order traversal follows the order**:
        D B E A C
- 