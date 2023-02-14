<?php

namespace App\Http\Controllers;

use App\Http\Resources\UserResource;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use PHPOpenSourceSaver\JWTAuth\Facades\JWTAuth;
use PHPOpenSourceSaver\JWTAuth\Exceptions\JWTException;
use PHPOpenSourceSaver\JWTAuth\Exceptions\TokenExpiredException;
use PHPOpenSourceSaver\JWTAuth\Exceptions\TokenInvalidException;

class UserController extends Controller
{
    /**
     * @OA\Post(
     *      path="/api/register",
     *      summary="User register endpoint",
     *      description="Endpoint to register new user",
     *      tags={"Register"},
     *
     *      @OA\RequestBody(
     *          @OA\MediaType(
     *              mediaType="application/json",
     *              @OA\Schema(
     *                  type="object",
     *
     *                  @OA\Property(
     *                     property="nik",
     *                     description="Identity number",
     *                     example="3574526883729839",
     *                     type="string"
     *                 ),
     *                 @OA\Property(
     *                     property="role",
     *                     description="User role",
     *                     example="admin",
     *                     type="string"
     *                 ),
     *                 @OA\Property(
     *                     property="password",
     *                     description="User password",
     *                     example="password",
     *                     type="string"
     *                 )
     *              )
     *          )
     *      ),
     *      @OA\Response(
     *          response=200,
     *          description="OK",
     *          @OA\MediaType(
     *              mediaType="application/json",
     *              @OA\Schema(
     *                  type="object",
     *                  @OA\Property(
     *                     property="success",
     *                     description="Response status",
     *                     example=true,
     *                     type="bool"
     *                 ),
     *                 @OA\Property(
     *                     property="message",
     *                     description="Response message",
     *                     example="Data user berhasil ditambahkan!",
     *                     type="string"
     *                 ),
     *                 @OA\Property(
     *                      property="data",
     *                      type="object",
     *                      @OA\Property(
     *                          property="nik",
     *                          description="User identity number",
     *                          example="3574526883729839",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="role",
     *                          description="User role",
     *                          example="admin",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="created_at",
     *                          description="Created user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="updated_at",
     *                          description="Updated user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="id",
     *                          description="Generated user id",
     *                          example=3,
     *                          type="int"
     *                      ),
     *                 )
     *              )
     *          )
     *      ),
     * )
     */
    public function register(Request $request)
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

    /**
     * @OA\Post(
     *      path="/api/login",
     *      summary="User login endpoint",
     *      description="Endpoint to login and genrate token",
     *      tags={"Login"},
     *
     *      @OA\RequestBody(
     *          @OA\MediaType(
     *              mediaType="application/json",
     *              @OA\Schema(
     *                  type="object",
     *
     *                  @OA\Property(
     *                     property="nik",
     *                     description="Identity number",
     *                     example="3574526883729838",
     *                     type="string"
     *                 ),
     *                 @OA\Property(
     *                     property="password",
     *                     description="User password",
     *                     example="password",
     *                     type="string"
     *                 )
     *              )
     *          )
     *      ),
     *      @OA\Response(
     *          response=200,
     *          description="OK",
     *          @OA\MediaType(
     *              mediaType="application/json",
     *              @OA\Schema(
     *                  type="object",
     *                  @OA\Property(
     *                     property="success",
     *                     description="Response status",
     *                     example=true,
     *                     type="bool"
     *                 ),
     *                 @OA\Property(
     *                      property="user",
     *                      type="object",
     *                      @OA\Property(
     *                          property="id",
     *                          description="Generated user id",
     *                          example=3,
     *                          type="int"
     *                      ),
     *                      @OA\Property(
     *                          property="nik",
     *                          description="User identity number",
     *                          example="3574526883729839",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="role",
     *                          description="User role",
     *                          example="admin",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="created_at",
     *                          description="Created user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="updated_at",
     *                          description="Updated user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      ),
     *                 ),
     *                 @OA\Property(
     *                     property="token",
     *                     description="Generated JWT Token",
     *                     example="xxxxxxxxxxxxxxxxxxxxxxxxx",
     *                     type="string"
     *                 )
     *              )
     *          )
     *      ),
     * )
     */
    public function login(Request $request)
    {
        //set validation
        $validator = Validator::make($request->all(), [
            'nik'     => 'required',
            'password'  => 'required'
        ]);

        //if validation fails
        if ($validator->fails()) {
            return response()->json($validator->errors(), 422);
        }

        //get credentials from request
        $credentials = $request->only('nik', 'password');

        //if auth failed
        if(!$token = auth()->guard('api')->attempt($credentials)) {
            return response()->json([
                'success' => false,
                'message' => 'NIK atau Password Anda salah'
            ], 401);
        }

        //if auth success
        return response()->json([
            'success' => true,
            'user'    => auth()->guard('api')->user(),
            'token'   => $token
        ], 200);
    }

    /**
     * @OA\Post(
     *      path="/api/verify",
     *      summary="Token verify endpoint",
     *      description="Endpoint to display private claims data",
     *      tags={"Verify"},
     *      security={
     *          {"bearerAuth": {}}
     *      },
     *      @OA\Response(
     *          response=200,
     *          description="OK",
     *          @OA\MediaType(
     *              mediaType="application/json",
     *              @OA\Schema(
     *                  type="object",
     *                  @OA\Property(
     *                     property="success",
     *                     description="Response status",
     *                     example=true,
     *                     type="bool"
     *                 ),
     *                 @OA\Property(
     *                     property="message",
     *                     description="Reponse message",
     *                     example="Token valid!",
     *                     type="string"
     *                 ),
     *                 @OA\Property(
     *                      property="user",
     *                      type="object",
     *                      @OA\Property(
     *                          property="id",
     *                          description="Generated user id",
     *                          example=3,
     *                          type="int"
     *                      ),
     *                      @OA\Property(
     *                          property="nik",
     *                          description="User identity number",
     *                          example="3574526883729839",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="role",
     *                          description="User role",
     *                          example="admin",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="created_at",
     *                          description="Created user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      ),
     *                      @OA\Property(
     *                          property="updated_at",
     *                          description="Updated user time",
     *                          example="2023-02-14T20:06:57.000000Z",
     *                          type="string"
     *                      )
     *                 )
     *              )
     *          )
     *      ),
     * )
     */
    public function verify()
    {
        //return response JSON
        return response()->json([
            'success' => true,
            'message' => 'Token valid!',
            'user' => auth()->guard('api')->user()
        ]);
    }
}
