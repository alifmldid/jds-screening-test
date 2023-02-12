<?php

namespace App\Http\Controllers;

use App\Http\Resources\UserResource;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class UserController extends Controller
{
    public function store(Request $request)
    {
        //define validation rules
        $validator = Validator::make($request->all(), [
            'nik'     => 'required|numeric|digits:16',
            'role'     => 'required',
            'password'   => 'required',
        ]);

        //check if validation fails
        if ($validator->fails()) {
            return response()->json($validator->errors(), 422);
        }

        //create user
        $user = User::create([
            'nik'       => $request->nik,
            'role'      => $request->role,
            'password'  => bcrypt($request->password),
        ]);

        //return response
        return new UserResource(true, 'Data user berhasil ditambahkan!', $user);
    }
}
