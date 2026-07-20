import cv2
import json
import logging
from ultralytics import YOLO

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger("CVPipeline")

class SafetyInferencePipeline:
    def __init__(self, model_path: str, confidence: float = 0.5):
        # Load YOLO Model (falls back to CPU if GPU not available)
        logger.info(f"Loading safety model weights from {model_path}...")
        self.model = YOLO(model_path)
        self.confidence = confidence

    def start_stream_inference(self, rtsp_url: str):
        """Starts real-time frame capture and object detection on an RTSP stream."""
        cap = cv2.VideoCapture(rtsp_url)
        if not cap.isOpened():
            logger.error(f"Failed to open RTSP video stream: {rtsp_url}")
            return

        logger.info(f"Stream ingestion active for: {rtsp_url}")
        try:
            while True:
                ret, frame = cap.read()
                if not ret:
                    logger.warning("Blank frame received. Reconnecting stream...")
                    break

                # Run Inference
                results = self.model(frame, conf=self.confidence, verbose=False)
                
                # Process predictions
                for result in results:
                    boxes = result.boxes
                    for box in boxes:
                        class_id = int(box.cls[0])
                        label = self.model.names[class_id]
                        conf_val = float(box.conf[0])
                        xyxy = box.xyxy[0].tolist()
                        
                        logger.debug(f"Detected {label} ({conf_val:.2f}) at {xyxy}")
                        
                        # In production, check rules (e.g. person detected with no hard hat)
                        # and write json violations to Kafka:
                        # self.publish_alert(label, conf_val, xyxy)

                # Press 'q' to break stream local loop
                if cv2.waitKey(1) & 0xFF == ord('q'):
                    break
        finally:
            cap.release()
            cv2.destroyAllWindows()
            logger.info("Video stream ingestion stopped.")

    def publish_alert(self, detection_label: str, confidence: float, bounding_box: list):
        """Standard alert format schema mapping for downstream Kafka ingestion."""
        alert_payload = {
            "facility_id": "plant-ue1-core",
            "camera_id": "cam-gate-01",
            "hazard_type": "PPE_VIOLATION",
            "details": f"Missing safety gear: {detection_label}",
            "coordinates": bounding_box,
            "confidence": confidence
        }
        logger.info(f"EMITTING KAFKA ALERT: {json.dumps(alert_payload)}")

if __name__ == "__main__":
    # Standard local test loop representation
    pipeline = SafetyInferencePipeline(model_path="yolov8n.pt", confidence=0.4)
    logger.info("CV Pipeline template instantiated successfully.")
