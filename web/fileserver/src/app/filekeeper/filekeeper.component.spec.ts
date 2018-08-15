import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FilekeeperComponent } from './filekeeper.component';

describe('FilekeeperComponent', () => {
  let component: FilekeeperComponent;
  let fixture: ComponentFixture<FilekeeperComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FilekeeperComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FilekeeperComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
