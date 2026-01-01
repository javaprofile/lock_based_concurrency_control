# Lock-Based Concurrency Control
**Designing Effective Lock-Based Concurrency Control in Database Systems**

### Paper Information
- **Author(s)**: Vipul Kumar Bondugula
- **Published In** : International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
- **Publication Date** : Feb 2021
- **ISSN** : E-ISSN 2147-6799
- **DOI** : 
- **Impact Factor** : 

**__Abstract__**\
This paper addresses throughput degradation in distributed database systems caused by inefficiencies in lease-based locking mechanisms. It examines how time-bound locks, while essential for fault tolerance and deadlock prevention, can lead to increased contention and reduced performance under high concurrency. The study emphasizes the role of fundamental locking strategies, including shared and exclusive locks, pessimistic and optimistic locking, and two-phase locking, in ensuring data consistency across distributed environments such as Kubernetes and etcd. By analyzing lease duration and lock acquisition behavior, the proposed approach improves lock utilization and reduces unnecessary blocking. The paper highlights the need for efficient, lightweight lease-based locking mechanisms to enhance throughput and reliability in large-scale distributed systems.

## Key Contributions
- **Lease-Based Locking Optimization:**
  Analyzed throughput limitations in lease-based locking mechanisms used in distributed database systems under high concurrency.

- **Concurrency Control Enhancement:**
  Studied shared, exclusive, pessimistic, optimistic, and two-phase locking models to understand their impact on data consistency and lock contention.
  
- **Performance Evaluation:**
  Evaluated the effect of lease duration, lock contention, and expiration behavior on throughput and transaction latency in distributed environments.
  
- **Research Leadership:**
  Led the study and implementation focusing on improving throughput and reliability of distributed systems using optimized lease-based locking mechanisms.

## Relevance & Real-World Impact
- **Improved Throughput in Distributed Systems:**
  Enhanced system performance by reducing lock contention and blocking caused by inefficient lease-based locking under concurrent workloads.

- **Reliable Fault-Tolerant Locking:**
  Strengthened automatic lock recovery and deadlock prevention in distributed platforms such as Kubernetes and etcd through effective lease management.

- **System Performance Improvement:**
  Reduced resource contention and improved transaction progress in high-traffic environments by refining lease-based lock acquisition and release behavior.
  
- **Academic & Research Impact:**
    Supports ongoing research and educational initiatives in distributed databases, locking protocols, and concurrency control for scalable systems.

## Experimental Results (Summary)

  | Nodes | Lease Based (locks/sec) | Basic Lease Based (locks/sec) | Gain (%)   |
  |-------|-------------------------| ------------------------------| ----------------|
  | 3     |  92                     | 105                           | 14.13           |
  | 5     |  88                     |  98                           | 11.36           |
  | 7     |  83                     |  93                           | 12.05           |
  | 9     |  78                     |  88                           | 12.82           |
  | 11    |  73                     |  82                           | 12.33           |

## Citation
DESIGNING EFFECTIVE LOCK-BASED CONCURRENCY CONTROL IN DATABASE SYSTEMS
* Vipul Kumar Bondugula
* International Journal of Intelligent Systems and Applications in Engineering 
* E-ISSN 2147-6799 
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://ijisae.org/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/in/vipul-b-18468a19/ | **Email**: vipulreddy574@gmail.com



