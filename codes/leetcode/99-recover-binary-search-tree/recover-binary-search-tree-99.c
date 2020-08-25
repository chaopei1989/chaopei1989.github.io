/*99. 恢复二叉搜索树
题目链接: https://leetcode-cn.com/problems/recover-binary-search-tree/

二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？

提示: 先考虑中序遍历
1324 -> 123
14325 -> 12345
21 -> 12
*/
#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

/********************************************************************
 *
 * 			以下是 Morris 中序遍历方式
 *
 *******************************************************************/

/**
 * 找前驱节点
 */
struct TreeNode *findPreceed(struct TreeNode *p)
{
    if (p->left == NULL) {
        return NULL;
    }
    struct TreeNode *pl = p->left;
    while (pl->right && pl->right != p) {
        pl = pl->right;
    }
    return pl;
}

void reverse(struct TreeNode *left, struct TreeNode *right)
{
    printf("reverse: left=%d, right=%d", left->val, right->val);
    left->val ^= right->val;
    right->val ^= left->val;
    left->val ^= right->val;
}

void recoverTree(struct TreeNode *root)
{
    struct TreeNode *res_1 = NULL;
    struct TreeNode *res_1_right = NULL;
    struct TreeNode *res_2 = NULL;
    struct TreeNode *last = NULL;
    struct TreeNode *curr = root;
    while (curr) {
        struct TreeNode *preceed = findPreceed(curr);
        if (preceed != NULL) {
            printf("recoverTree: preceed=%d\n", preceed->val);
            if (preceed->right == NULL) {
                // 前驱节点右节点设置为curr
                preceed->right = curr;
                curr = curr->left;
            } else if (preceed->right == curr) {
                // 前驱节点已经设置了, 且就是当前节点
                preceed->right = NULL;
                printf("recoverTree: curr=%d\n", curr->val);
                // TODO: do with curr
                if (last != NULL) {
                    if (last->val > curr->val) {
                        if (res_1 == NULL) {
                            res_1 = last;
                            res_1_right = curr;
                        } else {
                            res_2 = curr;
                        }
                    }
                }
                if (res_1 != NULL && res_2 != NULL) {
                    goto result;
                }
                last = curr;
                curr = curr->right;
            }
        } else {
            // 没有左节点时, 前驱节点才会为空
            // TODO: do with curr
            printf("recoverTree: curr=%d\n", curr->val);
            if (last != NULL) {
                if (last->val > curr->val) {
                    if (res_1 == NULL) {
                        res_1 = last;
                        res_1_right = curr;
                    } else {
                        res_2 = curr;
                    }
                }
            }
            if (res_1 != NULL && res_2 != NULL) {
                goto result;
            }
            last = curr;
            curr = curr->right;
        }
    }

result:

    if (res_2 == NULL) {
        res_2 = res_1_right;
    }
    reverse(res_1, res_2);
}

/********************************************************************
 *
 * 			以下是测试代码
 *
 *******************************************************************/

struct TreeNode *build_from_array(const char **array, int len, int index)
{
    if (index >= len) {
        return NULL;
    }
    if (array[index] == NULL) {
        return NULL;
    }
    // dfs 构建对象
    struct TreeNode *curr = (struct TreeNode *) malloc(sizeof(struct TreeNode));
    curr->val = atoi(array[index]);
    struct TreeNode *left = build_from_array(array, len, 2 * index + 1);
    curr->left = left;
    struct TreeNode *right = build_from_array(array, len, 2 * index + 2);
    curr->right = right;

    return curr;
}

int main(int argc, char *argv[])
{
    // const char *test[] = {
    //     "146",
    //     "71", "-13",
    //     "55", NULL, "231", "399",
    //     "321", NULL, NULL, NULL, NULL, NULL, NULL, NULL, "-33"
    // };
    const char *test[] = {
        "10", "5", "15", "0", "8", "13", "20", "2", "-5", "6", "9", "12", "14", "18", "25"
    };
    size_t len = sizeof(test) / sizeof(char *);
    struct TreeNode *root = build_from_array(&test[0], len, 0);

    printf("start\n");
    recoverTree(root);

    return 0;
}
