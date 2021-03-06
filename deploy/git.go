package deploy

import (
	"bytes"
	//"errors"
	//"fmt"
	"log"
	//"github.com/megamsys/gulp/scm"
	"io"
)

// Clone runs a git clone to clone the app repository in an app.
func clone() ([]byte, error) {
	var buf bytes.Buffer
	/*path, err := repository.GetPath()
	if err != nil {
		return nil, fmt.Errorf("Megam is misconfigured: %s", err)
	}
	cmd := fmt.Sprintf("git clone %s %s --depth 1", repository.ReadOnlyURL(app.GetName()), path)
	err = p.ExecuteCommand(&buf, &buf, app, cmd)
	*/
	b := buf.Bytes()
	log.Printf(`"git clone" output: %s`, b)
	//return b, err
	return b, nil
}

// fetch runs a git fetch to update the code in the app.
//
// It works like Clone, fetching from the app remote repository.
func fetch() ([]byte, error) {
	var buf bytes.Buffer
	/*
	path, err := repository.GetPath()
	if err != nil {
		return nil, fmt.Errorf("Megam is misconfigured: %s", err)
	}
	cmd := fmt.Sprintf("cd %s && git fetch origin", path)
	err = p.ExecuteCommand(&buf, &buf, app, cmd)
	*/
	b := buf.Bytes()
	log.Printf(`"git fetch" output: %s`, b)
	//return b, err
	return b, nil
	
}

// checkout updates the Git repository of the app to the given version.
func checkout(version string) ([]byte, error) {
/*	var buf bytes.Buffer
	path, err := repository.GetPath()
	if err != nil {
		return nil, fmt.Errorf("Megam is misconfigured: %s", err)
	}
	cmd := fmt.Sprintf("cd %s && git checkout %s", path, version)
	if err := p.ExecuteCommand(&buf, &buf, app, cmd); err != nil {
		return buf.Bytes(), err
	}
	*/
	return nil, nil
}

func Git(objID string, w io.Writer) error {
	/*log.Write(w, []byte("\n ---> Megam receiving push\n"))
	log.Write(w, []byte("\n ---> Replicating the application repository across units\n"))
	out, err := clone(provisioner, app)
	if err != nil {
		out, err = fetch(provisioner, app)
	}
	if err != nil {
		msg := fmt.Sprintf("Got error while cloning/fetching repository: %s -- \n%s", err.Error(), string(out))
		log.Write(w, []byte(msg))
		return errors.New(msg)
	}
	out, err = checkout(provisioner, app, objID)
	if err != nil {
		msg := fmt.Sprintf("Failed to checkout Git repository: %s -- \n%s", err, string(out))
		log.Write(w, []byte(msg))
		return errors.New(msg)
	}
	log.Write(w, []byte("\n ---> Installing dependencies\n"))
	if err := provisioner.InstallDeps(app, w); err != nil {
		log.Write(w, []byte(err.Error()))
		return err
	}
	log.Write(w, []byte("\n ---> Restarting application\n"))
	if err := app.Restart(w); err != nil {
		log.Write(w, []byte(err.Error()))
		return err
	}
	return log.Write(w, []byte("\n ---> Deploy done!\n\n"))
	*/
	return nil
}
