<h1>Welcome to Notestool Manual</h1>

The Notestool is a database modification tool that allows to create new databeses and modify existing ones.
Databases are stored in separate <i>.txt</i> files.

### Usage

#### 1. Starting the program:
Open the program via terminal with the command <code>>./notestool myNotesExample</code> where <i>myNoteExample</i> refers to an existing database <i>.txt</i> file.
(Use <code>>./notestool help</code> for further help. )

#### 2. Operations
If Notestool doesn't find your specified database file, it will create a new database with that name for note storing.

In the initial run you are greeted after which you are presented with the main menu with 4 operations to choose from via terminal commands.


    Welcome to NotesTool!

    Select operation (1-4):
    1. Show Notes
    2. Add Note
    3. Remove Note
    4. Exit



After any succesful operation (ex: 4. Exit) you will be returned back to the main menu.

*If your input happened to be unacceptable, Notestool will display helpful error messages to guide you back on track.*

#### 2.1 Show Notes
Inputing "1" will list all the existing notes in the current database. Notes will be numbered. 
Example:

    001 - First note text (Tags: [Tech], Timestamp: 18 Aug 24 09:34 EEST)
    002 - Second note text (Tags: [Cooking], Timestamp: 18 Aug 24 09:34 EEST)
    ...etc

If no notes are present, you will be informed of that.

#### 2.2 Add Notes
Inputting "2" will prompt you to write the contents of your note. Additionally, you will be prompted to input a Tags to help distinguish the notes more easily. Each note will also have a timestamp indicating when it was created.

In case of an empty input, the Notes Tool will let the user know that an empty note was added.


    Welcome to NotesTool!

    Select operation (1-4):
    1. Show Notes
    2. Add Note
    3. Remove Note
    4. Exit

    
    Enter your choice: 2

    Enter the note text:
    Meeting notes

    Enter tags (comma-separated):
    work,meeting


Note added successfully with timestamp: 18 Aug 24 09:34 EEST

Full note:

001 - Meeting notes (Tags: [work meeting], Timestamp: 18 Aug 24 09:42 EEST)


#### 2.3 Remove Note
Inputing "3" will ask the user for the number of the note to be removed.
By inputing "0" you will cancel this operation.

#### 2.4 Exit

Inputing "4" will exit the program and you are returned to an empty terminal.
