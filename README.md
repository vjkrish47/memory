# memory
**Reducing Runtime Overhead in Distributed Congestion Monitoring Systems**

* Author: Vijaya Krishna Namala
* Published In : International Journal of Innovative Research in Engineering & Multidisciplinary Physical Sciences (IJIRMPS)
* Publication Date: May 2022
* E-ISSN: 2349-7300
* Impact Factor : 9.907

**Abstract:**
Distributed transaction systems rely on commit protocols, but conventional approaches introduce high latency due to coordination overhead and multiple communication rounds. As systems scale, these delays reduce concurrency and increase sensitivity to network conditions. This work explores protocol level optimizations to restructure coordination and improve decision propagation while maintaining correctness. Experimental evaluation across cluster sizes demonstrates that optimized commit design can significantly reduce latency and improve scalability in distributed environments.

**Key Contributions**
* **Memory Efficient Monitoring Framework Design:**\
Developed a distributed monitoring architecture that minimizes memory overhead by eliminating redundant telemetry storage across nodes.
* **Consolidated Telemetry Storage Mechanism:**\
Introduced a shared aggregation approach where telemetry data is centrally maintained, avoiding duplication of buffers, logs, and intermediate data across machines.
* **Reduced Runtime Memory Overhead Strategy:**\
Designed lightweight data collection techniques that limit local buffering and reduce unnecessary memory allocation during monitoring operations.
* **Scalability Evaluation Across Cluster Sizes:**\
Conducted experiments on clusters with 3, 5, 7, 9, and 11 nodes to analyze memory consumption behavior and validate improved scalability.
**Relevance & Real World Impact**
* **Reduced Memory Consumption :**\
The proposed approach significantly lowers memory usage by avoiding duplicated telemetry storage and optimizing buffer management in distributed monitoring systems.
* **Improved Resource Utilization :**\
Efficient memory handling ensures that more system resources remain available for application workloads, improving overall system performance.
* **Enhanced Scalability in Distributed Environments :**\
Controlled memory growth with increasing cluster size enables scalable monitoring without excessive resource overhead.
* **Stable System Performance :**\
Lower memory pressure reduces risks of garbage collection delays and paging, ensuring consistent and reliable monitoring behavior.
* **Academic and Practical Contribution :**\
Provides a structured approach for designing resource efficient monitoring systems, supporting further research and real world implementation in cloud and distributed infrastructures.

**Experimental Results (Summary)**:

  | Nodes | Conventional 2PC Protocol Lock Hold Time (ms) | Optimized Commit Transaction Protocol Lock Hold Time (ms) | Improvement (%) |
  |-------|-----------------------------------------------| ----------------------------------------------------------| ----------------|
  | 3     | 95                                            | 65                                                        | 31.58           |
  | 5     | 120                                           | 85                                                        | 29.17           |
  | 7     | 145                                           | 105                                                       | 27.59           |
  | 9     | 170                                           | 125                                                       | 26.47           |
  | 11    | 195                                           | 145                                                       | 25.64           |

**Citation** \
Enhanced Commit Protocols for Low Latency Distributed Transactions. \
Vijaya Krishna Namala \
International Journal of Innovative Research in Engineering & Multidisciplinary Physical Sciences \
E-ISSN- 2349-7300  \
**License** \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijirmps.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/vijaya-krishna-namala-a42b2958 | **Email**: vijaya.namala@gmail.com


