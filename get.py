import json

def collect_keys_types(json_object, result_dict):
    for key, value in json_object.items():
        # Record the type of the value associated with the current key
        result_dict[key] = type(value).__name__

        # If the value is a nested object, call the function recursively
        if type(value) is dict:
            collect_keys_types(value, result_dict)
        
        # If the value is an array of objects, loop through each object
        elif type(value) is list and all(type(item) is dict for item in value):
            for item in value:
                collect_keys_types(item, result_dict)

# Initialize an empty dictionary to store unique keys and their types
unique_keys_with_types = {}

# Read the JSON array from the file
with open('test.json', 'r') as f:
    json_array = json.load(f)

# Traverse the array and collect keys along with their types
for json_object in json_array:
    collect_keys_types(json_object, unique_keys_with_types)

# Print the dictionary of unique keys and their types
print("Dictionary of unique keys with their types:", unique_keys_with_types)
