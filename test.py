import sys
import json

def process_resume(file_path):

    result = {
        "file": file_path,
        "status": "Processed",
        "data": {"name": "John Doe", "skills": ["Python", "Go"]}
    }
    return result

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({"error": "No file path provided"}))
        sys.exit(1)

    file_path = sys.argv[1]
    result = process_resume(file_path)
    print(json.dumps(result))  # Output JSON to stdout
