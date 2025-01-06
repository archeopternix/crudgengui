# Tasks Documentation
Tasks are used to define the needed actions to generate code or copy files to a target directory. It determines whether to generate a single file, multiple files using 
templates or copy file content without modifying it, and writes the output accordingly. It handles errors and logs the results.

## Index
* [Copy File Task](#copy-file-task)
* [Single File Template Task](#Single-File-Template-Task)
* [Multiple File Template Task](#Multiple-File-Template-Task)

## Copy File Task

The `copy` task is a type of generation task that copies files from a source directory to a target directory. This task is used to duplicate files exactly as they are, without any modifications or templating.

### Structure

A `copy` task is defined in the following structure within the YAML configuration:

```yaml
kind: copy
source:
  - "path/to/source/file1"
  - "path/to/source/file2"
target: "path/to/target/directory"
```

### Fields
* kind: Specifies the type of task. For a copy task, this value should be copy.
* source: A list of file paths that should be copied to the target directory. Each path should point to a file that exists in the source directory.
* target: The directory where the source files will be copied to. The target path should be a valid directory path where the copied files will be placed.

### Example
Below is an example of a copy task defined in a YAML configuration:

```yaml
kind: copy
source:
  - "templates/header.html"
  - "templates/footer.html"
target: "output/static"
```
In this example, the files header.html and footer.html located in the templates directory will be copied to the output/static directory.

### Usage
When the GenerateModule function processes a copy task, it performs the following steps for each file in the source list:

* It combines the genpath and target paths to determine the destination directory.
* It creates the target directory if it does not already exist.
* It copies each file from the source list to the target directory.
* If a file with the same name already exists in the target directory, the existing file will be overwritten.

## Single File Template Task
The single file generation task uses a template to generate a single output file based on the entire application structure. This task is useful when you want to create a file that summarizes or utilizes the whole application's data.

### Structure
A single file generation task is defined in the following structure within the YAML configuration:

```YAML
kind: template
source:
  - "path/to/template1"
  - "path/to/template2"
target: "path/to/target/directory"
template: "template_name"
fileext: ".ext"
filename: "output_filename"
```

### Fields
* kind: Specifies the type of task. For a single file generation task, this value should be template.
* source: A list of template file paths that will be used for generation.
* target: The directory where the generated file will be placed.
* template: The name of the primary template used for generation.
* fileext: The extension of the generated file.
* filename: The name of the generated file without the extension. When this is set, the whole application will be passed to the template execution.

### Example
Below is an example of a single file generation task defined in a YAML configuration:

```YAML
kind: template
source:
  - "templates/app_summary.tmpl"
target: "output"
template: "summary"
fileext: ".txt"
filename: "app_summary"
```

In this example, the app_summary.tmpl template is used to generate a app_summary.txt file in the output directory.

## Multiple File Template Task
The multiple file generation task uses a template to generate multiple output files, each based on an entity within the application. This task is useful when you want to create a file for each entity, such as individual model or view files.

### Structure
A multiple file generation task is defined in the following structure within the YAML configuration:

```YAML
kind: template
source:
  - "path/to/template1"
  - "path/to/template2"
target: "path/to/target/directory"
template: "template_name"
fileext: ".ext"
```

### Fields
* kind: Specifies the type of task. For a multiple file generation task, this value should be template.
* source: A list of template file paths that will be used for generation.
* target: The directory where the generated files will be placed.
* template: The name of the primary template used for generation.
* fileext: The extension of the generated files.

### Example
Below is an example of a multiple file generation task defined in a YAML configuration:

```YAML
kind: template
source:
  - "templates/entity.tmpl"
target: "output/entities"
template: "entity"
fileext: ".go"
```

In this example, the entity.tmpl template is used to generate Go files for each entity in the output/entities directory.

### Usage
When the GenerateModule function processes a template task, it performs the following steps:

1. It combines the genpath and target paths to determine the destination directory.
2. It creates the target directory if it does not already exist.
3. It parses the template files specified in the source list.
4. For single file generation:
  * It creates a file with the specified filename and fileext.
  * It passes the entire application data to the template and writes the output to the file.
5. For multiple file generation:
  * It iterates over each entity in the application.
  * It creates a file for each entity with the entity name and fileext.
  * It passes the entity data to the template and writes the output to the file.

