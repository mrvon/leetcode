/**
 * Definition for undirected graph.
 * struct UndirectedGraphNode {
 *     int label;
 *     vector<UndirectedGraphNode *> neighbors;
 *     UndirectedGraphNode(int x) : label(x) {};
 * };
 */
class Solution {
    public:
        UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
            map<int, UndirectedGraphNode*> share;
            return clone(share, node);
        }

    private:
        UndirectedGraphNode *clone(map<int, UndirectedGraphNode*> &share, UndirectedGraphNode *node) {
            if (node == nullptr) {
                return nullptr;
            }
            auto newNode = share[node->label];
            if (newNode == nullptr) {
                newNode = new UndirectedGraphNode(node->label);
                share[node->label] = newNode;
                for (auto node : node->neighbors) {
                    newNode->neighbors.push_back(clone(share, node));
                }
            }
            return newNode;
        }
};
