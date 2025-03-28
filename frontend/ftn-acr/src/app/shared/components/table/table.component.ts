import { NgFor } from '@angular/common';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-table',
  imports: [],
  templateUrl: './table.component.html',
  styleUrl: './table.component.scss',
})
export class TableComponent implements OnInit {
  tableData = [
    { name: 'Nenad', surname: 'Radovic' },
    { name: 'Borko', surname: 'Radovic' },
  ];

  ngOnInit(): void {}
}
