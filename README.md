Small and extensible anomaly detector system.

- It can use both internal implementations in Go and external services through REST/gRPC.
- There are several services already present using models created with sklearn, torch and multiple algorithms provided by the pyod library.
List with all available algorithms implemented in pyod:https://pyod.readthedocs.io/en/latest/index.html

---

There are two components:
 - The primary golang library
 - The services that expose the detection models and algorithms

---

TODO:
- [ ] Add Elliptic Envelope Algorithm