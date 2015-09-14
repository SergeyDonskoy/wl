package wundergo

// Client represents the methods that the API supports.
type Client interface {
	User() (User, error)
	UpdateUser(user User) (User, error)
	Users() ([]User, error)
	UsersForListID(listID uint) ([]User, error)

	Lists() ([]List, error)
	List(listID uint) (List, error)
	CreateList(title string) (List, error)
	UpdateList(list List) (List, error)
	DeleteList(list List) error

	NotesForListID(listID uint) ([]Note, error)
	NotesForTaskID(taskID uint) ([]Note, error)
	Note(noteID uint) (Note, error)
	CreateNote(content string, taskID uint) (Note, error)
	UpdateNote(note Note) (Note, error)
	DeleteNote(note Note) error

	TasksForListID(listID uint) ([]Task, error)
	CompletedTasksForListID(listID uint, completed bool) ([]Task, error)
	Task(taskID uint) (Task, error)
	CreateTask(
		title string,
		listID uint,
		assigneeID uint,
		completed bool,
		recurrenceType string,
		recurrenceCount uint,
		dueDate string,
		starred bool,
	) (Task, error)
	UpdateTask(task Task) (Task, error)
	DeleteTask(task Task) error

	SubtasksForListID(listID uint) ([]Subtask, error)
	SubtasksForTaskID(taskID uint) ([]Subtask, error)
	CompletedSubtasksForListID(listID uint, completed bool) ([]Subtask, error)
	CompletedSubtasksForTaskID(taskID uint, completed bool) ([]Subtask, error)
	Subtask(subtaskID uint) (Subtask, error)
	CreateSubtask(
		title string,
		taskID uint,
		completed bool,
	) (Subtask, error)
	UpdateSubtask(subtask Subtask) (Subtask, error)
	DeleteSubtask(subtask Subtask) error

	RemindersForListID(listID uint) ([]Reminder, error)
	RemindersForTaskID(taskID uint) ([]Reminder, error)
	Reminder(reminderID uint) (Reminder, error)
	CreateReminder(
		date string,
		taskID uint,
		createdByDeviceUdid string,
	) (Reminder, error)
	UpdateReminder(reminder Reminder) (Reminder, error)
	DeleteReminder(reminder Reminder) error

	ListPositions() ([]Position, error)
	ListPosition(listPositionID uint) (Position, error)
	UpdateListPosition(listPosition Position) (Position, error)

	TaskPositionsForListID(listID uint) ([]Position, error)
	TaskPosition(taskPositionID uint) (Position, error)
	UpdateTaskPosition(taskPosition Position) (Position, error)

	SubtaskPositionsForListID(listID uint) ([]Position, error)
	SubtaskPositionsForTaskID(taskID uint) ([]Position, error)
	SubtaskPosition(subtaskPositionID uint) (Position, error)
	UpdateSubtaskPosition(subtaskPosition Position) (Position, error)

	Memberships() ([]Membership, error)
	MembershipsForListID(listID uint) ([]Membership, error)
	Membership(membershipID uint) (Membership, error)
	AddMemberToListViaUserID(userID uint, listID uint, muted bool) (Membership, error)
	AddMemberToListViaEmailAddress(emailAddress string, listID uint, muted bool) (Membership, error)
	RejectInvite(membership Membership) error
	RemoveMemberFromList(membership Membership) error
	AcceptMember(membership Membership) (Membership, error)

	TaskCommentsForListID(listID uint) ([]TaskComment, error)
	TaskCommentsForTaskID(taskID uint) ([]TaskComment, error)
	CreateTaskComment(text string, taskID uint) (TaskComment, error)
	TaskComment(taskCommentID uint) (TaskComment, error)
	DeleteTaskComment(taskComment TaskComment) error

	AvatarURL(userID uint, size int, fallback bool) (string, error)

	CreateWebhook(listID uint, url string, processorType string, configuration string) (Webhook, error)
	DeleteWebhook(webhook Webhook) error
	WebhooksForListID(listID uint) ([]Webhook, error)

	Folders() ([]Folder, error)
	CreateFolder(title string, listIDs []uint) (Folder, error)
	Folder(folderID uint) (Folder, error)
	UpdateFolder(folder Folder) (Folder, error)
	DeleteFolder(folder Folder) error
	FolderRevisions() ([]FolderRevision, error)
	DeleteAllFolders() error
}
