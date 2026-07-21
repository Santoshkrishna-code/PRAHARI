class RecursiveCharacterTextSplitter:
    """Recursively splits a text into segments based on size and overlap parameters."""

    def __init__(self, chunk_size: int = 800, chunk_overlap: int = 100) -> None:
        self.chunk_size = chunk_size
        self.chunk_overlap = chunk_overlap
        self.separators = ["\n\n", "\n", " ", ""]

    def split_text(self, text: str) -> list[str]:
        """Runs the segmenting algorithm returning text chunk segments."""
        if not text:
            return []

        chunks = []
        start = 0
        text_len = len(text)

        while start < text_len:
            # Determine end position
            end = min(start + self.chunk_size, text_len)
            
            if end < text_len:
                # Find separator backwards to avoid breaking words mid-sentence
                found_sep = False
                for sep in self.separators[:-1]:
                    # Search backwards for separator within last 150 chars
                    last_sep_idx = text.rfind(sep, max(start, end - 150), end)
                    if last_sep_idx != -1:
                        end = last_sep_idx + len(sep)
                        found_sep = True
                        break
            
            chunk = text[start:end].strip()
            if chunk:
                chunks.append(chunk)

            # Slide window forward, accounting for overlap
            new_start = end - self.chunk_overlap
            if new_start <= start:
                start = end
            else:
                start = new_start

        return chunks
