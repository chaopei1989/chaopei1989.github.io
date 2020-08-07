/*114. 二叉树展开为链表
题目链接: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树
    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

注意: 需要保持前序遍历的顺序
*/
import java.lang.String;

public class FlattenBinaryTreeToLinkedList114 {
    public static void main(Stirng[] argv) {

    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) {
        this.val = val;
    }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {

    public void flatten(TreeNode root) {
        // 有空用例
        if (root == null) {
            return;
        }
        TreeNode origRight = root.right;
        if (root.left != null) {
            // 左分支改为链表
            flatten(root.left);
            // 将左分支挂到右边, 左分支记得置空
            root.right = root.left;
            root.left = null;
            // 挂原右分支
            if (origRight != null) {
                TreeNode tmp = root;
                while (tmp.right != null) {
                    tmp = tmp.right;
                }
                tmp.right = origRight;
            }
        }
        // 将原右分支继续改为链表
        if (origRight != null) {
            flatten(origRight);
        }
    }

}
