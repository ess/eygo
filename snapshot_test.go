package eygo

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"testing"
)

func TestNewSnapshotService(t *testing.T) {
	driver := NewMockDriver()
	service := NewSnapshotService(driver)

	t.Run("it is configured with the given driver", func(t *testing.T) {
		if service.Driver != driver {
			t.Errorf("Expected the service to use the given driver")
		}
	})
}

func TestSnapshotService_ForEnvironment(t *testing.T) {
	environment := &Environment{ID: 1, Name: "Environment 1"}
	driver := NewMockDriver()
	service := NewSnapshotService(driver)

	t.Run("when there are matching snapshots", func(t *testing.T) {
		snapshot1 := &Snapshot{ID: 1}
		snapshot2 := &Snapshot{ID: 2}
		snapshot3 := &Snapshot{ID: 3}

		stubEnvironmentSnapshots(driver, environment, snapshot1, snapshot2, snapshot3)

		all := service.ForEnvironment(environment, nil)

		t.Run("it contains all matching snapshots", func(t *testing.T) {
			snapshots := []*Snapshot{snapshot1, snapshot2, snapshot3}

			if len(all) != len(snapshots) {
				t.Errorf("Expected %d snapshots, got %d", len(snapshots), len(all))
			}

			for _, snapshot := range snapshots {
				found := false

				for _, other := range all {
					if snapshot.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Snapshot %d was not present", snapshot.ID)
				}
			}
		})

	})

	t.Run("when there are no matching snapshots", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForEnvironment(environment, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 snapshots, got")
			}
		})

	})

}

func TestSnapshotService_ForServer(t *testing.T) {
	server := &Server{ID: 1}
	driver := NewMockDriver()
	service := NewSnapshotService(driver)

	t.Run("when there are matching snapshots", func(t *testing.T) {
		snapshot1 := &Snapshot{ID: 1}
		snapshot2 := &Snapshot{ID: 2}
		snapshot3 := &Snapshot{ID: 3}

		stubServerSnapshots(driver, server, snapshot1, snapshot2, snapshot3)

		all := service.ForServer(server, nil)

		t.Run("it contains all matching snapshots", func(t *testing.T) {
			snapshots := []*Snapshot{snapshot1, snapshot2, snapshot3}

			if len(all) != len(snapshots) {
				t.Errorf("Expected %d snapshots, got %d", len(snapshots), len(all))
			}

			for _, snapshot := range snapshots {
				found := false

				for _, other := range all {
					if snapshot.ID == other.ID {
						found = true
					}
				}

				if !found {
					t.Errorf("Snapshot %d was not present", snapshot.ID)
				}
			}
		})

	})

	t.Run("when there are no matching snapshots", func(t *testing.T) {
		driver.Reset()

		t.Run("it is empty", func(t *testing.T) {
			all := service.ForServer(server, nil)

			if len(all) != 0 {
				t.Errorf("Expected 0 snapshots, got")
			}
		})

	})

}

func TestSnapshotService_Find(t *testing.T) {
	driver := NewMockDriver()
	service := NewSnapshotService(driver)
	snapshot := &Snapshot{ID: 1}
	stubSnapshot(driver, snapshot)

	t.Run("for a known snapshot", func(t *testing.T) {
		result, err := service.Find("1")

		t.Run("it is the snapshoted snapshot", func(t *testing.T) {
			if result.ID != snapshot.ID {
				t.Errorf("Expected snapshot 1, got snapshot %d", result.ID)
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error")
			}
		})
	})

	t.Run("for an unknown snapshot", func(t *testing.T) {
		result, err := service.Find("2")

		t.Run("it returns no snapshot", func(t *testing.T) {
			if result != nil {
				t.Errorf("Expected no snapshot, got snapshot %d", result.ID)
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}

func stubSnapshots(driver *MockDriver, snapshots ...*Snapshot) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Snapshots []*Snapshot `json:"snapshots,omitempty"`
	}{Snapshots: snapshots}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "snapshots", Response{Pages: pages})
	}
}

func stubEnvironmentSnapshots(driver *MockDriver, environment *Environment, snapshots ...*Snapshot) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Snapshots []*Snapshot `json:"snapshots,omitempty"`
	}{Snapshots: snapshots}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "environments/"+strconv.Itoa(environment.ID)+"/snapshots", Response{Pages: pages})
	}
}

func stubServerSnapshots(driver *MockDriver, server *Server, snapshots ...*Snapshot) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Snapshots []*Snapshot `json:"snapshots,omitempty"`
	}{Snapshots: snapshots}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "servers/"+strconv.Itoa(server.ID)+"/snapshots", Response{Pages: pages})
	}
}

func stubSnapshot(driver *MockDriver, snapshot *Snapshot) {
	pages := make([][]byte, 0)

	wrapper := struct {
		Snapshot *Snapshot `json:"snapshot,omitempty"`
	}{Snapshot: snapshot}

	if encoded, err := json.Marshal(&wrapper); err == nil {
		pages = append(pages, encoded)
		driver.AddResponse("get", "snapshots/"+strconv.Itoa(snapshot.ID), Response{Pages: pages})
	}
}
