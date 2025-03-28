import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TableComponent } from './shared/components/table/table.component';

@Component({
  selector: 'app-root',
  imports: [TableComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  title = 'ftn-acr';
}
