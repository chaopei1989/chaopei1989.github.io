struct Node {
    struct Node* last;
    struct Node* next;
    int content;
};

int* deckRevealedIncreasing(int* deck, int deckSize, int* returnSize)
{
    *returnSize = deckSize;
    int* res = (int*)malloc(sizeof(int) * deckSize);
    struct Node* head = NULL;
    struct Node* tail = NULL;
    // O(n^2) sort incread head~tail
    for (int i = 0; i < deckSize; i++) {
        for (int j = deckSize - 1; j > i; j--) {
            if (deck[j - 1] > deck[j]) {
                deck[j] = deck[j] ^ deck[j - 1];
                deck[j - 1] = deck[j] ^ deck[j - 1];
                deck[j] = deck[j] ^ deck[j - 1];
            }
        }
        struct Node* node = (struct Node*)malloc(sizeof(struct Node));
        node->content = deck[i];
        if (head) {
            if (head != tail) {
                node->last = tail;
                node->next = NULL;
                tail->next = node;
                tail = node;
            } else {
                head->next = node;
                node->last = head;
                node->next = NULL;
                tail = node;
            }
        } else {
            node->last = node->next = NULL;
            head = tail = node;
        }
    }
    // back trace
    struct Node* backHead = NULL;
    struct Node* backTail = NULL;
    for (struct Node* i = tail; i != NULL; i = i->last) {
        struct Node* node = (struct Node*)malloc(sizeof(struct Node));
        node->content = i->content;
        if (!backHead) {
            node->last = node->next = NULL;
            backHead = backTail = node;
        } else {
            // insert head
            node->last = NULL;
            node->next = backHead;
            backHead->last = node;
            backHead = node;
            // trans tail to head
            if (i->last) {
                struct Node* trans = backTail;
                backTail = trans->last;
                backTail->next = NULL;
                if (backTail == backHead) {
                    backTail->last = trans;
                }
                trans->last = NULL;
                trans->next = backHead;
                backHead->last = trans;
                backHead = trans;
            }
        }
    }
    // print
    struct Node* tmpNode = backHead;
    for (int i = 0; i < deckSize; i++) {
        res[i] = tmpNode->content;
        tmpNode = tmpNode->next;
    }
    return res;
}
