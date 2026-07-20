# Computer Vision Pipeline Template

This template establishes folder layouts and development guidelines for all Computer Vision analytics modules (e.g. PPE violation, fire/smoke monitoring, intrusion zones) in the **PRAHARI** platform.

## Folder Structure

* **`datasets/`**: Data directories split into `train/`, `val/`, and `test/` subfolders.
* **`annotations/`**: YOLO annotation coordinate text files (`.txt`).
* **`configs/`**: Setup variables (`model_config.yaml` specifying model class labels).
* **`training/`**: Scripts containing model training logic.
* **`validation/`**: Scripts evaluating accuracy metrics (mAP, Precision, Recall).
* **`inference/`**: Real-time object recognition pipeline loop.
* **`weights/`**: Folder to store target weights (`.pt`, `.onnx`, `.engine`).
* **`scripts/`**: Helper files (e.g. data splitting, annotations converter).
* **`tests/`**: Unit/Integration test suites.

## How to Initialize a New Pipeline

1. Copy the contents of this folder into `computer-vision/<new-pipeline-name>`.
2. Configure labels and weights locations inside `configs/model_config.yaml`.
3. Add dataset splits under `datasets/` (or map shared AWS S3 paths).
4. Run testing script:
   ```bash
   python inference/pipeline.py
   ```
