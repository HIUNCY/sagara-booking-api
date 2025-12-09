package service

import (
    "errors"
    "testing"

    "github.com/HIUNCY/sagara-booking-api/internal/core/domain"
    "github.com/HIUNCY/sagara-booking-api/internal/core/port"
)

type mockFieldRepo struct {
    byID map[uint]*domain.Field
    createErr error
    updateErr error
    deleteErr error
}

func (m *mockFieldRepo) Create(f *domain.Field) error {
    if m.createErr != nil {
        return m.createErr
    }
    if m.byID == nil {
        m.byID = map[uint]*domain.Field{}
    }
    f.ID = uint(len(m.byID) + 1)
    m.byID[f.ID] = f
    return nil
}

func (m *mockFieldRepo) GetAll() ([]domain.Field, error) {
    res := make([]domain.Field, 0, len(m.byID))
    for _, f := range m.byID {
        res = append(res, *f)
    }
    return res, nil
}

func (m *mockFieldRepo) GetByID(id uint) (*domain.Field, error) {
    if f, ok := m.byID[id]; ok {
        return f, nil
    }
    return nil, errors.New("not found")
}

func (m *mockFieldRepo) Update(f *domain.Field) error {
    if m.updateErr != nil {
        return m.updateErr
    }
    if _, ok := m.byID[f.ID]; !ok {
        return errors.New("not found")
    }
    m.byID[f.ID] = f
    return nil
}

func (m *mockFieldRepo) Delete(id uint) error {
    if m.deleteErr != nil {
        return m.deleteErr
    }
    if _, ok := m.byID[id]; !ok {
        return errors.New("not found")
    }
    delete(m.byID, id)
    return nil
}

func TestFieldService_CRUD(t *testing.T) {
    repo := &mockFieldRepo{}
    svc := NewFieldService(repo)

    // create
    if err := svc.CreateField(&port.CreateFieldRequest{Name: "A", PricePerHour: 10, Location: "L"}); err != nil {
        t.Fatalf("create error: %v", err)
    }
    all, _ := svc.GetAllFields()
    if len(all) != 1 {
        t.Fatalf("expected 1, got %d", len(all))
    }
    f, _ := svc.GetFieldByID(all[0].ID)
    if f.Name != "A" {
        t.Fatalf("unexpected field: %+v", f)
    }

    // update
    if err := svc.UpdateField(f.ID, &port.CreateFieldRequest{Name: "B", PricePerHour: 20, Location: "X"}); err != nil {
        t.Fatalf("update error: %v", err)
    }
    f2, _ := svc.GetFieldByID(f.ID)
    if f2.Name != "B" || f2.PricePerHour != 20 || f2.Location != "X" {
        t.Fatalf("update not applied: %+v", f2)
    }

    // delete
    if err := svc.DeleteField(f.ID); err != nil {
        t.Fatalf("delete error: %v", err)
    }
    all2, _ := svc.GetAllFields()
    if len(all2) != 0 {
        t.Fatalf("expected 0 after delete, got %d", len(all2))
    }
}
