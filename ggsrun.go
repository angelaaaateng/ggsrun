// Package main (ggsrun.go) :
// This file is included all commands and options.
package main

import (
	"os"

	"github.com/urfave/cli"
)

// main : main function
func main() {
	app := cli.NewApp()
	app.Name = appname
	app.Author = "Tanaike [ https://github.com/tanaikech/ggsrun ] "
	app.Email = "tanaike@hotmail.com"
	app.Usage = "Executes Google Apps Script (GAS) on Google and Feeds Back Results."
	app.Version = "1.5.0"
	app.Commands = []cli.Command{
		{
			Name:        "exe1",
			Aliases:     []string{"e1"},
			Usage:       "Updates project and Executes the function in the project.",
			Description: "In this mode, an access token is required.",
			Action:      exeAPIWithout,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scriptid, i",
					Usage: "Script ID of project on Google Drive",
				},
				cli.StringFlag{
					Name:  "scriptfile, s",
					Usage: "GAS file (.gs, .gas, .js, .coffee) on local PC",
				},
				cli.StringFlag{
					Name:  "function, f",
					Usage: "Function name which is executed. Default is '" + deffuncwithout + "'.",
				},
				cli.StringFlag{
					Name:  "value, v",
					Usage: "Give a value to the function which is executed.",
				},
				cli.BoolFlag{
					Name:  "backup, b",
					Usage: "Backup project with script ID you set as a file.",
				},
				cli.BoolFlag{
					Name:  "onlyresult, r",
					Usage: "Display only 'result' in JSON results",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "exe2",
			Aliases:     []string{"e2"},
			Usage:       "Uploads GAS and Executes the script using Execution API.",
			Description: "In this mode, an access token is required.",
			Action:      exeAPIWith,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "scriptid, i",
					Usage: "Script ID of project on Google Drive",
				},
				cli.StringFlag{
					Name:  "scriptfile, s",
					Usage: "GAS file (.gs, .gas, .js, .coffee) on local PC",
				},
				cli.StringFlag{
					Name:  "function, f",
					Usage: "Function name of server for executing GAS. Default is '" + deffuncserv + "'. If you change the server, use this.",
				},
				cli.StringFlag{
					Name:  "value, v",
					Usage: "Give a value to the function of GAS script which is executed.",
				},
				cli.StringFlag{
					Name:  "stringscript, ss",
					Usage: "GAS script as strings.",
				},
				cli.BoolFlag{
					Name:  "foldertree, t",
					Usage: "Display a folder tree on Google Drive as an array.",
				},
				cli.BoolFlag{
					Name:  "convert, conv",
					Usage: "[Experiment] Download file using byte slice data. Use with '-v [File ID]'.",
				},
				cli.BoolFlag{
					Name:  "log, l",
					Usage: "Record access log.",
				},
				cli.BoolFlag{
					Name:  "onlyresult, r",
					Usage: "Display only 'result' in JSON results",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "webapps",
			Aliases:     []string{"w"},
			Usage:       "Uploads GAS and Executes the script without OAuth using Web Apps.",
			Description: "In this mode, an access token is NOT required.",
			Action:      webAppsWith,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url, u",
					Usage: "URL for using Web Apps.",
				},
				cli.StringFlag{
					Name:  "scriptfile, s",
					Usage: "GAS file (.gs, .gas, .js, .coffee) on local PC",
				},
				cli.StringFlag{
					Name:  "value, v",
					Usage: "Give a value to the function of GAS script which is executed.",
				},
				cli.StringFlag{
					Name:  "password, p",
					Usage: "Password to use Web Apps (if you have set)",
				},
				cli.BoolFlag{
					Name:  "log, l",
					Usage: "Not record access log. No this option means 'Record log'.",
				},
				cli.BoolFlag{
					Name:  "onlyresult, r",
					Usage: "Display only 'result' in JSON results",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "download",
			Aliases:     []string{"d"},
			Usage:       "Downloads files from Google Drive.",
			Description: "In this mode, an access token is required.",
			Action:      downloadFiles,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "fileid, i",
					Usage: "File ID on Google Drive. Using file ID, you can download all files except for bound scripts.",
				},
				cli.StringFlag{
					Name:  "filename, f",
					Usage: "File Name on Google Drive",
				},
				cli.StringFlag{
					Name:  "extension, e",
					Usage: "Extension (File format of downloaded file)",
				},
				cli.BoolFlag{
					Name:  "rawdata, r",
					Usage: "Save a project with GAS scripts as raw data (JSON data).",
				},
				cli.BoolFlag{
					Name:  "zip, z",
					Usage: "Create a zip file including all scripts of a GAS project. Please use this for downloading a GAS project.",
				},
				cli.StringFlag{
					Name:  "deletefile",
					Usage: "Value is file ID. This can delete a file using a file ID on Google Drive.",
				},
				cli.BoolFlag{
					Name:  "overwrite, o",
					Usage: "When filename of downloading file is existing in directory at local PC, overwrite it. At default, it is not overwritten.",
				},
				cli.BoolFlag{
					Name:  "skip, s",
					Usage: "When filename of downloading file is existing in directory at local PC, skip it. At default, it is not overwritten.",
				},
				cli.BoolFlag{
					Name:  "showfilelist, l",
					Usage: "When files from a folder are retrieved, file list is returned using this option. When this is used, files are not downloaded.",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "upload",
			Aliases:     []string{"u"},
			Usage:       "Uploads files to Google Drive.",
			Description: "In this mode, an access token is required.",
			Action:      uploadFiles,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filename, f",
					Usage: "File Name on local PC. Please input files you want to upload.",
				},
				cli.StringFlag{
					Name:  "parentfolderid, p",
					Usage: "Folder ID of parent folder on Google Drive",
				},
				cli.StringFlag{
					Name:  "parentid, pid",
					Usage: "File ID of Google Docs (Spreadsheet, Document, Slide, Form) for creating container bound-script.",
				},
				cli.StringFlag{
					Name:  "timezone, tz",
					Usage: "Time zone of project. Please use this together with creating new project. When new project is created by API, time zone doesn't become the local time zone. (This might be a bug.) So please input this.",
				},
				cli.StringFlag{
					Name:  "projectname, pn",
					Usage: "Upload several GAS scripts as a project.",
				},
				cli.StringFlag{
					Name:  "googledocname, gn",
					Usage: "Filename of Google Docs which is created.",
				},
				cli.StringFlag{
					Name:  "projecttype, pt",
					Usage: "You can select where it creates a new project. Please input 'spreadsheet', 'document', 'slide' and 'form'. When you select one of them, new project is created as a bound script. If this option is not used, new project is created as a standalone script. This is a default.",
					Value: "standalone",
				},
				cli.Int64Flag{
					Name:  "chunksize, chunk",
					Usage: "You can also set the maximum chunk size for the resumable upload. This unit is MB.",
					Value: 100,
				},
				cli.BoolFlag{
					Name:  "noconvert, nc",
					Usage: "If you don't want to convert file to Google Apps format.",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "updateproject",
			Aliases:     []string{"ud"},
			Usage:       "Updates project on Google Drive.",
			Description: "In this mode, an access token is required.",
			Action:      updateProject,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filename, f",
					Usage: "File name. It's source files for updating. When you set files which are not in the project, the files are added to the project. When you set files which are in the project, the files are overwritten to the files with same filename.",
				},
				cli.BoolFlag{
					Name:  "deletefiles",
					Usage: "When you use this bool flag, projectid and filename, they are removed from the project.",
				},
				cli.StringFlag{
					Name:  "projectid, p",
					Usage: "ID of existing project. It's a destination project for updating.",
				},
				cli.BoolFlag{
					Name:  "backup, b",
					Usage: "Backup project with project ID you set as a file.",
				},
				cli.BoolFlag{
					Name:  "rearrange, r",
					Usage: "Interactively rearrange scripts in project using your terminal.",
				},
				cli.StringFlag{
					Name:  "rearrangewithfile, rf",
					Usage: "Rearrange scripts in project using a file.",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "revisionfiles",
			Aliases:     []string{"r"},
			Usage:       "Retrieves revision list and files.",
			Description: "In this mode, an access token is required.",
			Action:      revisionFiles,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "fileid, i",
					Usage: "Value is file ID. Display revision ID list.",
				},
				cli.StringFlag{
					Name:  "download, d",
					Usage: "Value is revision ID. Download revision file using it and file ID.",
				},
				cli.StringFlag{
					Name:  "createversion, cv",
					Usage: "Create new version of GAS project. Please input the description of version as string.",
				},
				cli.StringFlag{
					Name:  "extension, e",
					Usage: "Extension (File format of downloaded file)",
				},
				cli.BoolFlag{
					Name:  "rawdata, r",
					Usage: "Save a project with GAS scripts as raw data (JSON data).",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "filelist",
			Aliases:     []string{"ls"},
			Usage:       "Outputs a file list on Google Drive.",
			Description: "In this mode, an access token is required.",
			Action:      showFileList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "searchbyname, sn",
					Usage: "Search file using File Name. Output File ID.",
				},
				cli.StringFlag{
					Name:  "searchbyid, si",
					Usage: "Search file using File ID. Output File Name.",
				},
				cli.BoolFlag{
					Name:  "stdout, s",
					Usage: "Output all file list to standard output.",
				},
				cli.BoolFlag{
					Name:  "file, f",
					Usage: "Output all file list to a JSON file.",
				},
				cli.BoolFlag{
					Name:  "jsonparser, j",
					Usage: "Display results by JSON parser",
				},
			},
		},
		{
			Name:        "auth",
			Usage:       "Retrieve access and refresh tokens. If you changed scopes, please use this.",
			Description: "In this mode, 'client_secret.json' and Scopes are required.",
			Action:      reAuth,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Usage: "Port number of temporal web server for retrieving authorization code.",
					Value: 8080,
				},
			},
		},
	}
	app.CommandNotFound = commandNotFound
	app.Run(os.Args)
}
