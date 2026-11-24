# Count Good Nodes in Binary Tree

> NeetCode 150 · Trees — DFS threading the path's running maximum downward; a node at or above it is good and raises the bar for its subtree. Go, stdlib only, tested.

## Problem

A node is *good* when no node on the path from the root down to it is greater. Count the good nodes (the root always qualifies).

## Approach

The definition compares each node against its ancestors — but ancestors need not be revisited if their *maximum* rides along as a recursion parameter. At each node: at-or-above the running max means good, and the max updates for everything below. This "thread the path property down as an argument" move is the same one validate-BST (dsa-135) uses with bounds and path-sum problems use with a running total — the lightest form of tree DP, where the parent hands the child exactly the summary it needs. Seeding with min-int keeps negative-valued roots good; `>=` (not `>`) lets equal values qualify — both pinned by tests.

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
