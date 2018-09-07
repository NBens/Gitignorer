# Gitignorer
Generate .gitignore files easily for multiple languages

# Usage:

1. New gitignore: ``` ./gitignorer -create Java ```
2. Multiple gitignores: ``` ./gitignorer -create C++,Python,Java,Go (No spaces between commas) ```
3. Update Gitignore files: ``` ./gitignorer -update ```
4. List languages: ``` ./gitignorer -list ```
5. Create a template: ``` ./gitignorer -create-template C++,Python,Java -name NameOfYourTemplate (No spaces between commas) ```
6. List templates: ``` ./gitignorer -list-templates ```
7. Use a template: ``` ./gitignorer -use-template TemplateName ```

# Templates:

Gitignorer allows you to "merge" different languages' .gitignore files into one, which is called a **Template** 

Templates are saved in Gitignorer's data directory(./gitignorer_data/Templates) 

You can edit templates using any text editor, and add whatever files you want to be ignored (Make sure they are saved with a name like Templatename.Template.gitignore)

To use a template, run the command: ./gitignorer -use-template TemplateName
