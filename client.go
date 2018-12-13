package wl

import "time"

// Client represents the methods that the API supports.
type Client interface {
	User() (User, error)
	UpdateUser(user User) (User, error)
	Users() ([]User, error)
	UsersForListID(listID uint64) ([]User, error)

	Lists() ([]List, error)
	List(listID uint64) (List, error)
	CreateList(title string) (List, error)
	UpdateList(list List) (List, error)
	DeleteList(list List) error
	DeleteAllLists() error
	Inbox() (List, error)

	Notes() ([]Note, error)
	NotesForListID(listID uint64) ([]Note, error)
	NotesForTaskID(taskID uint64) ([]Note, error)
	Note(noteID uint64) (Note, error)
	CreateNote(content string, taskID uint64) (Note, error)
	UpdateNote(note Note) (Note, error)
	DeleteNote(note Note) error

	Tasks() ([]Task, error)
	CompletedTasks(completed bool) ([]Task, error)
	TasksForListID(listID uint64) ([]Task, error)
	CompletedTasksForListID(listID uint64, completed bool) ([]Task, error)
	Task(taskID uint64) (Task, error)
	CreateTask(
		title string,
		listID uint64,
		assigneeID uint64,
		completed bool,
		recurrenceType string,
		recurrenceCount uint64,
		dueDate time.Time,
		starred bool,
	) (Task, error)
	UpdateTask(task Task) (Task, error)
	DeleteTask(task Task) error
	DeleteAllTasks() error

	Subtasks() ([]Subtask, error)
	SubtasksForListID(listID uint64) ([]Subtask, error)
	SubtasksForTaskID(taskID uint64) ([]Subtask, error)
	CompletedSubtasks(completed bool) ([]Subtask, error)
	CompletedSubtasksForListID(listID uint64, completed bool) ([]Subtask, error)
	CompletedSubtasksForTaskID(taskID uint64, completed bool) ([]Subtask, error)
	Subtask(subtaskID uint64) (Subtask, error)
	CreateSubtask(
		title string,
		taskID uint64,
		completed bool,
	) (Subtask, error)
	UpdateSubtask(subtask Subtask) (Subtask, error)
	DeleteSubtask(subtask Subtask) error

	Reminders() ([]Reminder, error)
	RemindersForListID(listID uint64) ([]Reminder, error)
	RemindersForTaskID(taskID uint64) ([]Reminder, error)
	Reminder(reminderID uint64) (Reminder, error)
	CreateReminder(
		date string,
		taskID uint64,
		createdByDeviceUdid string,
	) (Reminder, error)
	UpdateReminder(reminder Reminder) (Reminder, error)
	DeleteReminder(reminder Reminder) error

	ListPositions() ([]Position, error)
	ListPosition(listPositionID uint64) (Position, error)
	UpdateListPosition(listPosition Position) (Position, error)

	TaskPositions() ([]Position, error)
	TaskPositionsForListID(listID uint64) ([]Position, error)
	TaskPosition(taskPositionID uint64) (Position, error)
	UpdateTaskPosition(taskPosition Position) (Position, error)

	SubtaskPositions() ([]Position, error)
	SubtaskPositionsForListID(listID uint64) ([]Position, error)
	SubtaskPositionsForTaskID(taskID uint64) ([]Position, error)
	SubtaskPosition(subtaskPositionID uint64) (Position, error)
	UpdateSubtaskPosition(subtaskPosition Position) (Position, error)

	Memberships() ([]Membership, error)
	MembershipsForListID(listID uint64) ([]Membership, error)
	Membership(membershipID uint64) (Membership, error)
	AddMemberToListViaUserID(userID uint64, listID uint64, muted bool) (Membership, error)
	AddMemberToListViaEmailAddress(emailAddress string, listID uint64, muted bool) (Membership, error)
	RejectInvite(membership Membership) error
	RemoveMemberFromList(membership Membership) error
	AcceptMember(membership Membership) (Membership, error)

	TaskComments() ([]TaskComment, error)
	TaskCommentsForListID(listID uint64) ([]TaskComment, error)
	TaskCommentsForTaskID(taskID uint64) ([]TaskComment, error)
	CreateTaskComment(text string, taskID uint64) (TaskComment, error)
	TaskComment(taskCommentID uint64) (TaskComment, error)
	DeleteTaskComment(taskComment TaskComment) error

	AvatarURL(userID uint64, size int, fallback bool) (string, error)

	Webhooks() ([]Webhook, error)
	WebhooksForListID(listID uint64) ([]Webhook, error)
	Webhook(webhookID uint64) (Webhook, error)
	CreateWebhook(listID uint64, url string, processorType string, configuration string) (Webhook, error)
	DeleteWebhook(webhook Webhook) error

	Folders() ([]Folder, error)
	CreateFolder(title string, listIDs []uint64) (Folder, error)
	Folder(folderID uint64) (Folder, error)
	UpdateFolder(folder Folder) (Folder, error)
	DeleteFolder(folder Folder) error
	FolderRevisions() ([]FolderRevision, error)
	DeleteAllFolders() error

	UploadFile(
		localFilePath string,
		remoteFileName string,
		contentType string,
		md5sum string,
	) (Upload, error)

	Files() ([]File, error)
	FilesForTaskID(taskID uint64) ([]File, error)
	FilesForListID(listID uint64) ([]File, error)
	File(fileID uint64) (File, error)
	CreateFile(uploadID uint64, taskID uint64) (File, error)
	DestroyFile(file File) error
	FilePreview(fileID uint64, platform string, size string) (FilePreview, error)

	Root() (Root, error)
}
