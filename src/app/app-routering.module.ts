import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UsersComponent} from './components/users';

const routes: Routes = [
  {path: 'accounts', component: UsersComponent},
  {path: '', redirectTo: '/accounts', pathMatch: 'full'},
  {path: '**', redirectTo: '/accounts'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule {}
