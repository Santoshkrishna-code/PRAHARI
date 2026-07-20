from langgraph.checkpoint.memory import MemorySaver

def get_in_memory_checkpointer() -> MemorySaver:
    """Returns a simple in-memory checkpointer for tracking state history.
    
    Replace with PostgresSaver/RedisSaver in production for persistent clustering.
    """
    return MemorySaver()
