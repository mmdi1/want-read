// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {db} from '../models';

export function GetBook():Promise<db.Book>;

export function GetBookshelf():Promise<any>;

export function Lock():Promise<void>;

export function ReloadPage(arg1:string,arg2:number):Promise<string>;

export function RemoveBook(arg1:string):Promise<boolean>;

export function SelectFile():Promise<string>;

export function TryLock():Promise<boolean>;

export function Unlock():Promise<void>;
