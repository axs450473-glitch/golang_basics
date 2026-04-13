

def load_instructions_file(filepath: str) -> str:
    """
    Reads the content of a text file and returns it as a string.
    
    Args:
        filepath (str): The relative or absolute path to the text file.
        
    Returns:
        str: The content of the file, or an error message if it fails.
    """
    try:
        # 'r' mode opens the file for reading. 
        # specifying utf-8 encoding prevents issues with special characters.
        with open(filepath, 'r', encoding='utf-8') as file:
            content = file.read()
            return content
            
    except FileNotFoundError:
        return f"Error: The file was not found at '{filepath}'. Please check the path."
    except PermissionError:
        return f"Error: You do not have permission to read the file at '{filepath}'."
    except Exception as e:
        return f"An unexpected error occurred while reading the file: {e}"